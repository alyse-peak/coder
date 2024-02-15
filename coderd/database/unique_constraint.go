// Code generated by scripts/dbgen/main.go. DO NOT EDIT.
package database

// UniqueConstraint represents a named unique constraint on a table.
type UniqueConstraint string

// UniqueConstraint enums.
const (
	UniqueAgentStatsPkey                                    UniqueConstraint = "agent_stats_pkey"                                         // ALTER TABLE ONLY workspace_agent_stats ADD CONSTRAINT agent_stats_pkey PRIMARY KEY (id);
	UniqueAPIKeysPkey                                       UniqueConstraint = "api_keys_pkey"                                            // ALTER TABLE ONLY api_keys ADD CONSTRAINT api_keys_pkey PRIMARY KEY (id);
	UniqueAuditLogsPkey                                     UniqueConstraint = "audit_logs_pkey"                                          // ALTER TABLE ONLY audit_logs ADD CONSTRAINT audit_logs_pkey PRIMARY KEY (id);
	UniqueDbcryptKeysActiveKeyDigestKey                     UniqueConstraint = "dbcrypt_keys_active_key_digest_key"                       // ALTER TABLE ONLY dbcrypt_keys ADD CONSTRAINT dbcrypt_keys_active_key_digest_key UNIQUE (active_key_digest);
	UniqueDbcryptKeysPkey                                   UniqueConstraint = "dbcrypt_keys_pkey"                                        // ALTER TABLE ONLY dbcrypt_keys ADD CONSTRAINT dbcrypt_keys_pkey PRIMARY KEY (number);
	UniqueDbcryptKeysRevokedKeyDigestKey                    UniqueConstraint = "dbcrypt_keys_revoked_key_digest_key"                      // ALTER TABLE ONLY dbcrypt_keys ADD CONSTRAINT dbcrypt_keys_revoked_key_digest_key UNIQUE (revoked_key_digest);
	UniqueFilesHashCreatedByKey                             UniqueConstraint = "files_hash_created_by_key"                                // ALTER TABLE ONLY files ADD CONSTRAINT files_hash_created_by_key UNIQUE (hash, created_by);
	UniqueFilesPkey                                         UniqueConstraint = "files_pkey"                                               // ALTER TABLE ONLY files ADD CONSTRAINT files_pkey PRIMARY KEY (id);
	UniqueGitAuthLinksProviderIDUserIDKey                   UniqueConstraint = "git_auth_links_provider_id_user_id_key"                   // ALTER TABLE ONLY external_auth_links ADD CONSTRAINT git_auth_links_provider_id_user_id_key UNIQUE (provider_id, user_id);
	UniqueGitSSHKeysPkey                                    UniqueConstraint = "gitsshkeys_pkey"                                          // ALTER TABLE ONLY gitsshkeys ADD CONSTRAINT gitsshkeys_pkey PRIMARY KEY (user_id);
	UniqueGroupMembersUserIDGroupIDKey                      UniqueConstraint = "group_members_user_id_group_id_key"                       // ALTER TABLE ONLY group_members ADD CONSTRAINT group_members_user_id_group_id_key UNIQUE (user_id, group_id);
	UniqueGroupsNameOrganizationIDKey                       UniqueConstraint = "groups_name_organization_id_key"                          // ALTER TABLE ONLY groups ADD CONSTRAINT groups_name_organization_id_key UNIQUE (name, organization_id);
	UniqueGroupsPkey                                        UniqueConstraint = "groups_pkey"                                              // ALTER TABLE ONLY groups ADD CONSTRAINT groups_pkey PRIMARY KEY (id);
	UniqueJfrogXrayScansPkey                                UniqueConstraint = "jfrog_xray_scans_pkey"                                    // ALTER TABLE ONLY jfrog_xray_scans ADD CONSTRAINT jfrog_xray_scans_pkey PRIMARY KEY (agent_id, workspace_id);
	UniqueLicensesJWTKey                                    UniqueConstraint = "licenses_jwt_key"                                         // ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_jwt_key UNIQUE (jwt);
	UniqueLicensesPkey                                      UniqueConstraint = "licenses_pkey"                                            // ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);
	UniqueOauth2ProviderAppSecretsAppIDHashedSecretKey      UniqueConstraint = "oauth2_provider_app_secrets_app_id_hashed_secret_key"     // ALTER TABLE ONLY oauth2_provider_app_secrets ADD CONSTRAINT oauth2_provider_app_secrets_app_id_hashed_secret_key UNIQUE (app_id, hashed_secret);
	UniqueOauth2ProviderAppSecretsPkey                      UniqueConstraint = "oauth2_provider_app_secrets_pkey"                         // ALTER TABLE ONLY oauth2_provider_app_secrets ADD CONSTRAINT oauth2_provider_app_secrets_pkey PRIMARY KEY (id);
	UniqueOauth2ProviderAppsNameKey                         UniqueConstraint = "oauth2_provider_apps_name_key"                            // ALTER TABLE ONLY oauth2_provider_apps ADD CONSTRAINT oauth2_provider_apps_name_key UNIQUE (name);
	UniqueOauth2ProviderAppsPkey                            UniqueConstraint = "oauth2_provider_apps_pkey"                                // ALTER TABLE ONLY oauth2_provider_apps ADD CONSTRAINT oauth2_provider_apps_pkey PRIMARY KEY (id);
	UniqueOrganizationMembersPkey                           UniqueConstraint = "organization_members_pkey"                                // ALTER TABLE ONLY organization_members ADD CONSTRAINT organization_members_pkey PRIMARY KEY (organization_id, user_id);
	UniqueOrganizationsPkey                                 UniqueConstraint = "organizations_pkey"                                       // ALTER TABLE ONLY organizations ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);
	UniqueParameterSchemasJobIDNameKey                      UniqueConstraint = "parameter_schemas_job_id_name_key"                        // ALTER TABLE ONLY parameter_schemas ADD CONSTRAINT parameter_schemas_job_id_name_key UNIQUE (job_id, name);
	UniqueParameterSchemasPkey                              UniqueConstraint = "parameter_schemas_pkey"                                   // ALTER TABLE ONLY parameter_schemas ADD CONSTRAINT parameter_schemas_pkey PRIMARY KEY (id);
	UniqueParameterValuesPkey                               UniqueConstraint = "parameter_values_pkey"                                    // ALTER TABLE ONLY parameter_values ADD CONSTRAINT parameter_values_pkey PRIMARY KEY (id);
	UniqueParameterValuesScopeIDNameKey                     UniqueConstraint = "parameter_values_scope_id_name_key"                       // ALTER TABLE ONLY parameter_values ADD CONSTRAINT parameter_values_scope_id_name_key UNIQUE (scope_id, name);
	UniqueProvisionerDaemonsPkey                            UniqueConstraint = "provisioner_daemons_pkey"                                 // ALTER TABLE ONLY provisioner_daemons ADD CONSTRAINT provisioner_daemons_pkey PRIMARY KEY (id);
	UniqueProvisionerJobLogsPkey                            UniqueConstraint = "provisioner_job_logs_pkey"                                // ALTER TABLE ONLY provisioner_job_logs ADD CONSTRAINT provisioner_job_logs_pkey PRIMARY KEY (id);
	UniqueProvisionerJobsPkey                               UniqueConstraint = "provisioner_jobs_pkey"                                    // ALTER TABLE ONLY provisioner_jobs ADD CONSTRAINT provisioner_jobs_pkey PRIMARY KEY (id);
	UniqueSiteConfigsKeyKey                                 UniqueConstraint = "site_configs_key_key"                                     // ALTER TABLE ONLY site_configs ADD CONSTRAINT site_configs_key_key UNIQUE (key);
	UniqueTailnetAgentsPkey                                 UniqueConstraint = "tailnet_agents_pkey"                                      // ALTER TABLE ONLY tailnet_agents ADD CONSTRAINT tailnet_agents_pkey PRIMARY KEY (id, coordinator_id);
	UniqueTailnetClientSubscriptionsPkey                    UniqueConstraint = "tailnet_client_subscriptions_pkey"                        // ALTER TABLE ONLY tailnet_client_subscriptions ADD CONSTRAINT tailnet_client_subscriptions_pkey PRIMARY KEY (client_id, coordinator_id, agent_id);
	UniqueTailnetClientsPkey                                UniqueConstraint = "tailnet_clients_pkey"                                     // ALTER TABLE ONLY tailnet_clients ADD CONSTRAINT tailnet_clients_pkey PRIMARY KEY (id, coordinator_id);
	UniqueTailnetCoordinatorsPkey                           UniqueConstraint = "tailnet_coordinators_pkey"                                // ALTER TABLE ONLY tailnet_coordinators ADD CONSTRAINT tailnet_coordinators_pkey PRIMARY KEY (id);
	UniqueTailnetPeersPkey                                  UniqueConstraint = "tailnet_peers_pkey"                                       // ALTER TABLE ONLY tailnet_peers ADD CONSTRAINT tailnet_peers_pkey PRIMARY KEY (id, coordinator_id);
	UniqueTailnetTunnelsPkey                                UniqueConstraint = "tailnet_tunnels_pkey"                                     // ALTER TABLE ONLY tailnet_tunnels ADD CONSTRAINT tailnet_tunnels_pkey PRIMARY KEY (coordinator_id, src_id, dst_id);
	UniqueTemplateVersionParametersTemplateVersionIDNameKey UniqueConstraint = "template_version_parameters_template_version_id_name_key" // ALTER TABLE ONLY template_version_parameters ADD CONSTRAINT template_version_parameters_template_version_id_name_key UNIQUE (template_version_id, name);
	UniqueTemplateVersionVariablesTemplateVersionIDNameKey  UniqueConstraint = "template_version_variables_template_version_id_name_key"  // ALTER TABLE ONLY template_version_variables ADD CONSTRAINT template_version_variables_template_version_id_name_key UNIQUE (template_version_id, name);
	UniqueTemplateVersionsPkey                              UniqueConstraint = "template_versions_pkey"                                   // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_pkey PRIMARY KEY (id);
	UniqueTemplateVersionsTemplateIDNameKey                 UniqueConstraint = "template_versions_template_id_name_key"                   // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_template_id_name_key UNIQUE (template_id, name);
	UniqueTemplatesPkey                                     UniqueConstraint = "templates_pkey"                                           // ALTER TABLE ONLY templates ADD CONSTRAINT templates_pkey PRIMARY KEY (id);
	UniqueUserLinksPkey                                     UniqueConstraint = "user_links_pkey"                                          // ALTER TABLE ONLY user_links ADD CONSTRAINT user_links_pkey PRIMARY KEY (user_id, login_type);
	UniqueUsersPkey                                         UniqueConstraint = "users_pkey"                                               // ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
	UniqueWorkspaceAgentLogSourcesPkey                      UniqueConstraint = "workspace_agent_log_sources_pkey"                         // ALTER TABLE ONLY workspace_agent_log_sources ADD CONSTRAINT workspace_agent_log_sources_pkey PRIMARY KEY (workspace_agent_id, id);
	UniqueWorkspaceAgentMetadataPkey                        UniqueConstraint = "workspace_agent_metadata_pkey"                            // ALTER TABLE ONLY workspace_agent_metadata ADD CONSTRAINT workspace_agent_metadata_pkey PRIMARY KEY (workspace_agent_id, key);
	UniqueWorkspaceAgentPortSharePkey                       UniqueConstraint = "workspace_agent_port_share_pkey"                          // ALTER TABLE ONLY workspace_agent_port_share ADD CONSTRAINT workspace_agent_port_share_pkey PRIMARY KEY (workspace_id, agent_name, port);
	UniqueWorkspaceAgentStartupLogsPkey                     UniqueConstraint = "workspace_agent_startup_logs_pkey"                        // ALTER TABLE ONLY workspace_agent_logs ADD CONSTRAINT workspace_agent_startup_logs_pkey PRIMARY KEY (id);
	UniqueWorkspaceAgentsPkey                               UniqueConstraint = "workspace_agents_pkey"                                    // ALTER TABLE ONLY workspace_agents ADD CONSTRAINT workspace_agents_pkey PRIMARY KEY (id);
	UniqueWorkspaceAppStatsPkey                             UniqueConstraint = "workspace_app_stats_pkey"                                 // ALTER TABLE ONLY workspace_app_stats ADD CONSTRAINT workspace_app_stats_pkey PRIMARY KEY (id);
	UniqueWorkspaceAppStatsUserIDAgentIDSessionIDKey        UniqueConstraint = "workspace_app_stats_user_id_agent_id_session_id_key"      // ALTER TABLE ONLY workspace_app_stats ADD CONSTRAINT workspace_app_stats_user_id_agent_id_session_id_key UNIQUE (user_id, agent_id, session_id);
	UniqueWorkspaceAppsAgentIDSlugIndex                     UniqueConstraint = "workspace_apps_agent_id_slug_idx"                         // ALTER TABLE ONLY workspace_apps ADD CONSTRAINT workspace_apps_agent_id_slug_idx UNIQUE (agent_id, slug);
	UniqueWorkspaceAppsPkey                                 UniqueConstraint = "workspace_apps_pkey"                                      // ALTER TABLE ONLY workspace_apps ADD CONSTRAINT workspace_apps_pkey PRIMARY KEY (id);
	UniqueWorkspaceBuildParametersWorkspaceBuildIDNameKey   UniqueConstraint = "workspace_build_parameters_workspace_build_id_name_key"   // ALTER TABLE ONLY workspace_build_parameters ADD CONSTRAINT workspace_build_parameters_workspace_build_id_name_key UNIQUE (workspace_build_id, name);
	UniqueWorkspaceBuildsJobIDKey                           UniqueConstraint = "workspace_builds_job_id_key"                              // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_job_id_key UNIQUE (job_id);
	UniqueWorkspaceBuildsPkey                               UniqueConstraint = "workspace_builds_pkey"                                    // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_pkey PRIMARY KEY (id);
	UniqueWorkspaceBuildsWorkspaceIDBuildNumberKey          UniqueConstraint = "workspace_builds_workspace_id_build_number_key"           // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_workspace_id_build_number_key UNIQUE (workspace_id, build_number);
	UniqueWorkspaceProxiesPkey                              UniqueConstraint = "workspace_proxies_pkey"                                   // ALTER TABLE ONLY workspace_proxies ADD CONSTRAINT workspace_proxies_pkey PRIMARY KEY (id);
	UniqueWorkspaceProxiesRegionIDUnique                    UniqueConstraint = "workspace_proxies_region_id_unique"                       // ALTER TABLE ONLY workspace_proxies ADD CONSTRAINT workspace_proxies_region_id_unique UNIQUE (region_id);
	UniqueWorkspaceResourceMetadataName                     UniqueConstraint = "workspace_resource_metadata_name"                         // ALTER TABLE ONLY workspace_resource_metadata ADD CONSTRAINT workspace_resource_metadata_name UNIQUE (workspace_resource_id, key);
	UniqueWorkspaceResourceMetadataPkey                     UniqueConstraint = "workspace_resource_metadata_pkey"                         // ALTER TABLE ONLY workspace_resource_metadata ADD CONSTRAINT workspace_resource_metadata_pkey PRIMARY KEY (id);
	UniqueWorkspaceResourcesPkey                            UniqueConstraint = "workspace_resources_pkey"                                 // ALTER TABLE ONLY workspace_resources ADD CONSTRAINT workspace_resources_pkey PRIMARY KEY (id);
	UniqueWorkspacesPkey                                    UniqueConstraint = "workspaces_pkey"                                          // ALTER TABLE ONLY workspaces ADD CONSTRAINT workspaces_pkey PRIMARY KEY (id);
	UniqueIndexAPIKeyName                                   UniqueConstraint = "idx_api_key_name"                                         // CREATE UNIQUE INDEX idx_api_key_name ON api_keys USING btree (user_id, token_name) WHERE (login_type = 'token'::login_type);
	UniqueIndexOrganizationName                             UniqueConstraint = "idx_organization_name"                                    // CREATE UNIQUE INDEX idx_organization_name ON organizations USING btree (name);
	UniqueIndexOrganizationNameLower                        UniqueConstraint = "idx_organization_name_lower"                              // CREATE UNIQUE INDEX idx_organization_name_lower ON organizations USING btree (lower(name));
	UniqueIndexProvisionerDaemonsNameOwnerKey               UniqueConstraint = "idx_provisioner_daemons_name_owner_key"                   // CREATE UNIQUE INDEX idx_provisioner_daemons_name_owner_key ON provisioner_daemons USING btree (name, lower(COALESCE((tags ->> 'owner'::text), ''::text)));
	UniqueIndexUsersEmail                                   UniqueConstraint = "idx_users_email"                                          // CREATE UNIQUE INDEX idx_users_email ON users USING btree (email) WHERE (deleted = false);
	UniqueIndexUsersUsername                                UniqueConstraint = "idx_users_username"                                       // CREATE UNIQUE INDEX idx_users_username ON users USING btree (username) WHERE (deleted = false);
	UniqueOrganizationsSingleDefaultOrg                     UniqueConstraint = "organizations_single_default_org"                         // CREATE UNIQUE INDEX organizations_single_default_org ON organizations USING btree (is_default) WHERE (is_default = true);
	UniqueTemplatesOrganizationIDNameIndex                  UniqueConstraint = "templates_organization_id_name_idx"                       // CREATE UNIQUE INDEX templates_organization_id_name_idx ON templates USING btree (organization_id, lower((name)::text)) WHERE (deleted = false);
	UniqueUsersEmailLowerIndex                              UniqueConstraint = "users_email_lower_idx"                                    // CREATE UNIQUE INDEX users_email_lower_idx ON users USING btree (lower(email)) WHERE (deleted = false);
	UniqueUsersUsernameLowerIndex                           UniqueConstraint = "users_username_lower_idx"                                 // CREATE UNIQUE INDEX users_username_lower_idx ON users USING btree (lower(username)) WHERE (deleted = false);
	UniqueWorkspaceProxiesLowerNameIndex                    UniqueConstraint = "workspace_proxies_lower_name_idx"                         // CREATE UNIQUE INDEX workspace_proxies_lower_name_idx ON workspace_proxies USING btree (lower(name)) WHERE (deleted = false);
	UniqueWorkspacesOwnerIDLowerIndex                       UniqueConstraint = "workspaces_owner_id_lower_idx"                            // CREATE UNIQUE INDEX workspaces_owner_id_lower_idx ON workspaces USING btree (owner_id, lower((name)::text)) WHERE (deleted = false);
)
