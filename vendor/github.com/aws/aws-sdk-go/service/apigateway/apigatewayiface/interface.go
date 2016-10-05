// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package apigatewayiface provides an interface for the Amazon API Gateway.
package apigatewayiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

// APIGatewayAPI is the interface type for apigateway.APIGateway.
type APIGatewayAPI interface {
	CreateApiKeyRequest(*apigateway.CreateApiKeyInput) (*request.Request, *apigateway.ApiKey)

	CreateApiKey(*apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error)

	CreateAuthorizerRequest(*apigateway.CreateAuthorizerInput) (*request.Request, *apigateway.Authorizer)

	CreateAuthorizer(*apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error)

	CreateBasePathMappingRequest(*apigateway.CreateBasePathMappingInput) (*request.Request, *apigateway.BasePathMapping)

	CreateBasePathMapping(*apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error)

	CreateDeploymentRequest(*apigateway.CreateDeploymentInput) (*request.Request, *apigateway.Deployment)

	CreateDeployment(*apigateway.CreateDeploymentInput) (*apigateway.Deployment, error)

	CreateDomainNameRequest(*apigateway.CreateDomainNameInput) (*request.Request, *apigateway.DomainName)

	CreateDomainName(*apigateway.CreateDomainNameInput) (*apigateway.DomainName, error)

	CreateModelRequest(*apigateway.CreateModelInput) (*request.Request, *apigateway.Model)

	CreateModel(*apigateway.CreateModelInput) (*apigateway.Model, error)

	CreateResourceRequest(*apigateway.CreateResourceInput) (*request.Request, *apigateway.Resource)

	CreateResource(*apigateway.CreateResourceInput) (*apigateway.Resource, error)

	CreateRestApiRequest(*apigateway.CreateRestApiInput) (*request.Request, *apigateway.RestApi)

	CreateRestApi(*apigateway.CreateRestApiInput) (*apigateway.RestApi, error)

	CreateStageRequest(*apigateway.CreateStageInput) (*request.Request, *apigateway.Stage)

	CreateStage(*apigateway.CreateStageInput) (*apigateway.Stage, error)

	DeleteApiKeyRequest(*apigateway.DeleteApiKeyInput) (*request.Request, *apigateway.DeleteApiKeyOutput)

	DeleteApiKey(*apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error)

	DeleteAuthorizerRequest(*apigateway.DeleteAuthorizerInput) (*request.Request, *apigateway.DeleteAuthorizerOutput)

	DeleteAuthorizer(*apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error)

	DeleteBasePathMappingRequest(*apigateway.DeleteBasePathMappingInput) (*request.Request, *apigateway.DeleteBasePathMappingOutput)

	DeleteBasePathMapping(*apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error)

	DeleteClientCertificateRequest(*apigateway.DeleteClientCertificateInput) (*request.Request, *apigateway.DeleteClientCertificateOutput)

	DeleteClientCertificate(*apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error)

	DeleteDeploymentRequest(*apigateway.DeleteDeploymentInput) (*request.Request, *apigateway.DeleteDeploymentOutput)

	DeleteDeployment(*apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error)

	DeleteDomainNameRequest(*apigateway.DeleteDomainNameInput) (*request.Request, *apigateway.DeleteDomainNameOutput)

	DeleteDomainName(*apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error)

	DeleteIntegrationRequest(*apigateway.DeleteIntegrationInput) (*request.Request, *apigateway.DeleteIntegrationOutput)

	DeleteIntegration(*apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error)

	DeleteIntegrationResponseRequest(*apigateway.DeleteIntegrationResponseInput) (*request.Request, *apigateway.DeleteIntegrationResponseOutput)

	DeleteIntegrationResponse(*apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error)

	DeleteMethodRequest(*apigateway.DeleteMethodInput) (*request.Request, *apigateway.DeleteMethodOutput)

	DeleteMethod(*apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error)

	DeleteMethodResponseRequest(*apigateway.DeleteMethodResponseInput) (*request.Request, *apigateway.DeleteMethodResponseOutput)

	DeleteMethodResponse(*apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error)

	DeleteModelRequest(*apigateway.DeleteModelInput) (*request.Request, *apigateway.DeleteModelOutput)

	DeleteModel(*apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error)

	DeleteResourceRequest(*apigateway.DeleteResourceInput) (*request.Request, *apigateway.DeleteResourceOutput)

	DeleteResource(*apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error)

	DeleteRestApiRequest(*apigateway.DeleteRestApiInput) (*request.Request, *apigateway.DeleteRestApiOutput)

	DeleteRestApi(*apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error)

	DeleteStageRequest(*apigateway.DeleteStageInput) (*request.Request, *apigateway.DeleteStageOutput)

	DeleteStage(*apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error)

	FlushStageAuthorizersCacheRequest(*apigateway.FlushStageAuthorizersCacheInput) (*request.Request, *apigateway.FlushStageAuthorizersCacheOutput)

	FlushStageAuthorizersCache(*apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error)

	FlushStageCacheRequest(*apigateway.FlushStageCacheInput) (*request.Request, *apigateway.FlushStageCacheOutput)

	FlushStageCache(*apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error)

	GenerateClientCertificateRequest(*apigateway.GenerateClientCertificateInput) (*request.Request, *apigateway.ClientCertificate)

	GenerateClientCertificate(*apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error)

	GetAccountRequest(*apigateway.GetAccountInput) (*request.Request, *apigateway.Account)

	GetAccount(*apigateway.GetAccountInput) (*apigateway.Account, error)

	GetApiKeyRequest(*apigateway.GetApiKeyInput) (*request.Request, *apigateway.ApiKey)

	GetApiKey(*apigateway.GetApiKeyInput) (*apigateway.ApiKey, error)

	GetApiKeysRequest(*apigateway.GetApiKeysInput) (*request.Request, *apigateway.GetApiKeysOutput)

	GetApiKeys(*apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error)

	GetApiKeysPages(*apigateway.GetApiKeysInput, func(*apigateway.GetApiKeysOutput, bool) bool) error

	GetAuthorizerRequest(*apigateway.GetAuthorizerInput) (*request.Request, *apigateway.Authorizer)

	GetAuthorizer(*apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error)

	GetAuthorizersRequest(*apigateway.GetAuthorizersInput) (*request.Request, *apigateway.GetAuthorizersOutput)

	GetAuthorizers(*apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error)

	GetBasePathMappingRequest(*apigateway.GetBasePathMappingInput) (*request.Request, *apigateway.BasePathMapping)

	GetBasePathMapping(*apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error)

	GetBasePathMappingsRequest(*apigateway.GetBasePathMappingsInput) (*request.Request, *apigateway.GetBasePathMappingsOutput)

	GetBasePathMappings(*apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error)

	GetBasePathMappingsPages(*apigateway.GetBasePathMappingsInput, func(*apigateway.GetBasePathMappingsOutput, bool) bool) error

	GetClientCertificateRequest(*apigateway.GetClientCertificateInput) (*request.Request, *apigateway.ClientCertificate)

	GetClientCertificate(*apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error)

	GetClientCertificatesRequest(*apigateway.GetClientCertificatesInput) (*request.Request, *apigateway.GetClientCertificatesOutput)

	GetClientCertificates(*apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error)

	GetClientCertificatesPages(*apigateway.GetClientCertificatesInput, func(*apigateway.GetClientCertificatesOutput, bool) bool) error

	GetDeploymentRequest(*apigateway.GetDeploymentInput) (*request.Request, *apigateway.Deployment)

	GetDeployment(*apigateway.GetDeploymentInput) (*apigateway.Deployment, error)

	GetDeploymentsRequest(*apigateway.GetDeploymentsInput) (*request.Request, *apigateway.GetDeploymentsOutput)

	GetDeployments(*apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error)

	GetDeploymentsPages(*apigateway.GetDeploymentsInput, func(*apigateway.GetDeploymentsOutput, bool) bool) error

	GetDomainNameRequest(*apigateway.GetDomainNameInput) (*request.Request, *apigateway.DomainName)

	GetDomainName(*apigateway.GetDomainNameInput) (*apigateway.DomainName, error)

	GetDomainNamesRequest(*apigateway.GetDomainNamesInput) (*request.Request, *apigateway.GetDomainNamesOutput)

	GetDomainNames(*apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error)

	GetDomainNamesPages(*apigateway.GetDomainNamesInput, func(*apigateway.GetDomainNamesOutput, bool) bool) error

	GetExportRequest(*apigateway.GetExportInput) (*request.Request, *apigateway.GetExportOutput)

	GetExport(*apigateway.GetExportInput) (*apigateway.GetExportOutput, error)

	GetIntegrationRequest(*apigateway.GetIntegrationInput) (*request.Request, *apigateway.Integration)

	GetIntegration(*apigateway.GetIntegrationInput) (*apigateway.Integration, error)

	GetIntegrationResponseRequest(*apigateway.GetIntegrationResponseInput) (*request.Request, *apigateway.IntegrationResponse)

	GetIntegrationResponse(*apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error)

	GetMethodRequest(*apigateway.GetMethodInput) (*request.Request, *apigateway.Method)

	GetMethod(*apigateway.GetMethodInput) (*apigateway.Method, error)

	GetMethodResponseRequest(*apigateway.GetMethodResponseInput) (*request.Request, *apigateway.MethodResponse)

	GetMethodResponse(*apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error)

	GetModelRequest(*apigateway.GetModelInput) (*request.Request, *apigateway.Model)

	GetModel(*apigateway.GetModelInput) (*apigateway.Model, error)

	GetModelTemplateRequest(*apigateway.GetModelTemplateInput) (*request.Request, *apigateway.GetModelTemplateOutput)

	GetModelTemplate(*apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error)

	GetModelsRequest(*apigateway.GetModelsInput) (*request.Request, *apigateway.GetModelsOutput)

	GetModels(*apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error)

	GetModelsPages(*apigateway.GetModelsInput, func(*apigateway.GetModelsOutput, bool) bool) error

	GetResourceRequest(*apigateway.GetResourceInput) (*request.Request, *apigateway.Resource)

	GetResource(*apigateway.GetResourceInput) (*apigateway.Resource, error)

	GetResourcesRequest(*apigateway.GetResourcesInput) (*request.Request, *apigateway.GetResourcesOutput)

	GetResources(*apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error)

	GetResourcesPages(*apigateway.GetResourcesInput, func(*apigateway.GetResourcesOutput, bool) bool) error

	GetRestApiRequest(*apigateway.GetRestApiInput) (*request.Request, *apigateway.RestApi)

	GetRestApi(*apigateway.GetRestApiInput) (*apigateway.RestApi, error)

	GetRestApisRequest(*apigateway.GetRestApisInput) (*request.Request, *apigateway.GetRestApisOutput)

	GetRestApis(*apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error)

	GetRestApisPages(*apigateway.GetRestApisInput, func(*apigateway.GetRestApisOutput, bool) bool) error

	GetSdkRequest(*apigateway.GetSdkInput) (*request.Request, *apigateway.GetSdkOutput)

	GetSdk(*apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error)

	GetStageRequest(*apigateway.GetStageInput) (*request.Request, *apigateway.Stage)

	GetStage(*apigateway.GetStageInput) (*apigateway.Stage, error)

	GetStagesRequest(*apigateway.GetStagesInput) (*request.Request, *apigateway.GetStagesOutput)

	GetStages(*apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error)

	ImportRestApiRequest(*apigateway.ImportRestApiInput) (*request.Request, *apigateway.RestApi)

	ImportRestApi(*apigateway.ImportRestApiInput) (*apigateway.RestApi, error)

	PutIntegrationRequest(*apigateway.PutIntegrationInput) (*request.Request, *apigateway.Integration)

	PutIntegration(*apigateway.PutIntegrationInput) (*apigateway.Integration, error)

	PutIntegrationResponseRequest(*apigateway.PutIntegrationResponseInput) (*request.Request, *apigateway.IntegrationResponse)

	PutIntegrationResponse(*apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error)

	PutMethodRequest(*apigateway.PutMethodInput) (*request.Request, *apigateway.Method)

	PutMethod(*apigateway.PutMethodInput) (*apigateway.Method, error)

	PutMethodResponseRequest(*apigateway.PutMethodResponseInput) (*request.Request, *apigateway.MethodResponse)

	PutMethodResponse(*apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error)

	PutRestApiRequest(*apigateway.PutRestApiInput) (*request.Request, *apigateway.RestApi)

	PutRestApi(*apigateway.PutRestApiInput) (*apigateway.RestApi, error)

	TestInvokeAuthorizerRequest(*apigateway.TestInvokeAuthorizerInput) (*request.Request, *apigateway.TestInvokeAuthorizerOutput)

	TestInvokeAuthorizer(*apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error)

	TestInvokeMethodRequest(*apigateway.TestInvokeMethodInput) (*request.Request, *apigateway.TestInvokeMethodOutput)

	TestInvokeMethod(*apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error)

	UpdateAccountRequest(*apigateway.UpdateAccountInput) (*request.Request, *apigateway.Account)

	UpdateAccount(*apigateway.UpdateAccountInput) (*apigateway.Account, error)

	UpdateApiKeyRequest(*apigateway.UpdateApiKeyInput) (*request.Request, *apigateway.ApiKey)

	UpdateApiKey(*apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error)

	UpdateAuthorizerRequest(*apigateway.UpdateAuthorizerInput) (*request.Request, *apigateway.Authorizer)

	UpdateAuthorizer(*apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error)

	UpdateBasePathMappingRequest(*apigateway.UpdateBasePathMappingInput) (*request.Request, *apigateway.BasePathMapping)

	UpdateBasePathMapping(*apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error)

	UpdateClientCertificateRequest(*apigateway.UpdateClientCertificateInput) (*request.Request, *apigateway.ClientCertificate)

	UpdateClientCertificate(*apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error)

	UpdateDeploymentRequest(*apigateway.UpdateDeploymentInput) (*request.Request, *apigateway.Deployment)

	UpdateDeployment(*apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error)

	UpdateDomainNameRequest(*apigateway.UpdateDomainNameInput) (*request.Request, *apigateway.DomainName)

	UpdateDomainName(*apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error)

	UpdateIntegrationRequest(*apigateway.UpdateIntegrationInput) (*request.Request, *apigateway.Integration)

	UpdateIntegration(*apigateway.UpdateIntegrationInput) (*apigateway.Integration, error)

	UpdateIntegrationResponseRequest(*apigateway.UpdateIntegrationResponseInput) (*request.Request, *apigateway.IntegrationResponse)

	UpdateIntegrationResponse(*apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error)

	UpdateMethodRequest(*apigateway.UpdateMethodInput) (*request.Request, *apigateway.Method)

	UpdateMethod(*apigateway.UpdateMethodInput) (*apigateway.Method, error)

	UpdateMethodResponseRequest(*apigateway.UpdateMethodResponseInput) (*request.Request, *apigateway.MethodResponse)

	UpdateMethodResponse(*apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error)

	UpdateModelRequest(*apigateway.UpdateModelInput) (*request.Request, *apigateway.Model)

	UpdateModel(*apigateway.UpdateModelInput) (*apigateway.Model, error)

	UpdateResourceRequest(*apigateway.UpdateResourceInput) (*request.Request, *apigateway.Resource)

	UpdateResource(*apigateway.UpdateResourceInput) (*apigateway.Resource, error)

	UpdateRestApiRequest(*apigateway.UpdateRestApiInput) (*request.Request, *apigateway.RestApi)

	UpdateRestApi(*apigateway.UpdateRestApiInput) (*apigateway.RestApi, error)

	UpdateStageRequest(*apigateway.UpdateStageInput) (*request.Request, *apigateway.Stage)

	UpdateStage(*apigateway.UpdateStageInput) (*apigateway.Stage, error)
}

var _ APIGatewayAPI = (*apigateway.APIGateway)(nil)
