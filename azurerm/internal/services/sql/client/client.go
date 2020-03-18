package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql"
	v3sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	DatabasesClient                          *sql.DatabasesClient
	DatabaseThreatDetectionPoliciesClient    *sql.DatabaseThreatDetectionPoliciesClient
	ElasticPoolsClient                       *sql.ElasticPoolsClient
	FirewallRulesClient                      *sql.FirewallRulesClient
	FailoverGroupsClient                     *sql.FailoverGroupsClient
	ServersClient                            *sql.ServersClient
	ServerAzureADAdministratorsClient        *sql.ServerAzureADAdministratorsClient
	VirtualNetworkRulesClient                *sql.VirtualNetworkRulesClient
	ExtendedServerBlobAuditingPoliciesClient *sql.ExtendedServerBlobAuditingPoliciesClient
	BackupLongTermRetentionPoliciesClient    *v3sql.BackupLongTermRetentionPoliciesClient
	BackupShortTermRetentionPoliciesClient   *v3sql.BackupShortTermRetentionPoliciesClient
}

func NewClient(o *common.ClientOptions) *Client {
	// SQL Azure
	DatabasesClient := sql.NewDatabasesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DatabasesClient.Client, o.ResourceManagerAuthorizer)

	DatabaseThreatDetectionPoliciesClient := sql.NewDatabaseThreatDetectionPoliciesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DatabaseThreatDetectionPoliciesClient.Client, o.ResourceManagerAuthorizer)

	ElasticPoolsClient := sql.NewElasticPoolsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ElasticPoolsClient.Client, o.ResourceManagerAuthorizer)

	FailoverGroupsClient := sql.NewFailoverGroupsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&FailoverGroupsClient.Client, o.ResourceManagerAuthorizer)

	FirewallRulesClient := sql.NewFirewallRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&FirewallRulesClient.Client, o.ResourceManagerAuthorizer)

	ServersClient := sql.NewServersClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ServersClient.Client, o.ResourceManagerAuthorizer)

	ServerAzureADAdministratorsClient := sql.NewServerAzureADAdministratorsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ServerAzureADAdministratorsClient.Client, o.ResourceManagerAuthorizer)

	VirtualNetworkRulesClient := sql.NewVirtualNetworkRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&VirtualNetworkRulesClient.Client, o.ResourceManagerAuthorizer)

	ExtendedServerBlobAuditingPoliciesClient := sql.NewExtendedServerBlobAuditingPoliciesClient(o.SubscriptionId)
	o.ConfigureClient(&ExtendedServerBlobAuditingPoliciesClient.Client, o.ResourceManagerAuthorizer)

	BackupLongTermRetentionPoliciesClient := v3sql.NewBackupLongTermRetentionPoliciesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&BackupLongTermRetentionPoliciesClient.Client, o.ResourceManagerAuthorizer)

	BackupShortTermRetentionPoliciesClient := v3sql.NewBackupShortTermRetentionPoliciesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&BackupShortTermRetentionPoliciesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		DatabasesClient:                          &DatabasesClient,
		DatabaseThreatDetectionPoliciesClient:    &DatabaseThreatDetectionPoliciesClient,
		ElasticPoolsClient:                       &ElasticPoolsClient,
		FailoverGroupsClient:                     &FailoverGroupsClient,
		FirewallRulesClient:                      &FirewallRulesClient,
		ServersClient:                            &ServersClient,
		ServerAzureADAdministratorsClient:        &ServerAzureADAdministratorsClient,
		VirtualNetworkRulesClient:                &VirtualNetworkRulesClient,
		ExtendedServerBlobAuditingPoliciesClient: &ExtendedServerBlobAuditingPoliciesClient,
		BackupLongTermRetentionPoliciesClient:    &BackupLongTermRetentionPoliciesClient,
		BackupShortTermRetentionPoliciesClient:   &BackupShortTermRetentionPoliciesClient,
	}
}
