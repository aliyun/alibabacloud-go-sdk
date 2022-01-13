// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Environment struct {
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// Human-readable description of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The kind of the resource
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources. Cannot be updated.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specification of the desired behavior of the Environment
	Spec *EnvironmentSpec `json:"spec,omitempty" xml:"spec,omitempty" type:"Struct"`
	// Most recently observed status of the Environment. This data may not be up-to-date. Populated by the system. Read-only
	Status *EnvironmentStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// Main user ID of an Aliyun account
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Environment) String() string {
	return tea.Prettify(s)
}

func (s Environment) GoString() string {
	return s.String()
}

func (s *Environment) SetCreationTime(v string) *Environment {
	s.CreationTime = &v
	return s
}

func (s *Environment) SetDescription(v string) *Environment {
	s.Description = &v
	return s
}

func (s *Environment) SetGeneration(v int32) *Environment {
	s.Generation = &v
	return s
}

func (s *Environment) SetKind(v string) *Environment {
	s.Kind = &v
	return s
}

func (s *Environment) SetName(v string) *Environment {
	s.Name = &v
	return s
}

func (s *Environment) SetSpec(v *EnvironmentSpec) *Environment {
	s.Spec = v
	return s
}

func (s *Environment) SetStatus(v *EnvironmentStatus) *Environment {
	s.Status = v
	return s
}

func (s *Environment) SetUid(v string) *Environment {
	s.Uid = &v
	return s
}

type EnvironmentSpec struct {
	// A region ID at Aliyun. For example, "cn-hangzhou"
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ARN (Aliyun Resource Name) of the role designated by a user to allow the system to manage his Aliyun resources. If null, use roleArn of role AliyunFCDefaultRole.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The name of the template for the Environment
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// Variables for specified template
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
}

func (s EnvironmentSpec) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentSpec) GoString() string {
	return s.String()
}

func (s *EnvironmentSpec) SetRegion(v string) *EnvironmentSpec {
	s.Region = &v
	return s
}

func (s *EnvironmentSpec) SetRoleArn(v string) *EnvironmentSpec {
	s.RoleArn = &v
	return s
}

func (s *EnvironmentSpec) SetTemplate(v string) *EnvironmentSpec {
	s.Template = &v
	return s
}

func (s *EnvironmentSpec) SetTemplateVariables(v map[string]interface{}) *EnvironmentSpec {
	s.TemplateVariables = v
	return s
}

type EnvironmentStatus struct {
	// Credentials required for tasks
	Credentials *StsCredentials `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// A human-readable message indicating details about why the Environment is in this condition
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The most recent generation observed
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// Time when the last update of the status is observed
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// Details of current state of the Environment
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// A simple, high-level summary of where the Environment is in its lifecycle
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s EnvironmentStatus) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentStatus) GoString() string {
	return s.String()
}

func (s *EnvironmentStatus) SetCredentials(v *StsCredentials) *EnvironmentStatus {
	s.Credentials = v
	return s
}

func (s *EnvironmentStatus) SetMessage(v string) *EnvironmentStatus {
	s.Message = &v
	return s
}

func (s *EnvironmentStatus) SetObservedGeneration(v int32) *EnvironmentStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *EnvironmentStatus) SetObservedTime(v string) *EnvironmentStatus {
	s.ObservedTime = &v
	return s
}

func (s *EnvironmentStatus) SetOutput(v map[string]interface{}) *EnvironmentStatus {
	s.Output = v
	return s
}

func (s *EnvironmentStatus) SetPhase(v string) *EnvironmentStatus {
	s.Phase = &v
	return s
}

type Service struct {
	// A time representing the server time when this object was created. Clients may not set this value. Populated by the system. Read-only.
	CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// Human-readable description of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The kind of the resource
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources. Cannot be updated.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specification of the desired behavior of the Service
	Spec *ServiceSpec `json:"spec,omitempty" xml:"spec,omitempty" type:"Struct"`
	// Most recently observed status of the Service. This data may not be up-to-date. Populated by the system. Read-only
	Status *ServiceStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// Main user ID of an Aliyun account
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Service) String() string {
	return tea.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) SetCreationTime(v string) *Service {
	s.CreationTime = &v
	return s
}

func (s *Service) SetDescription(v string) *Service {
	s.Description = &v
	return s
}

func (s *Service) SetGeneration(v int32) *Service {
	s.Generation = &v
	return s
}

func (s *Service) SetKind(v string) *Service {
	s.Kind = &v
	return s
}

func (s *Service) SetName(v string) *Service {
	s.Name = &v
	return s
}

func (s *Service) SetSpec(v *ServiceSpec) *Service {
	s.Spec = v
	return s
}

func (s *Service) SetStatus(v *ServiceStatus) *Service {
	s.Status = v
	return s
}

func (s *Service) SetUid(v string) *Service {
	s.Uid = &v
	return s
}

type ServiceSpec struct {
	// The name of the associated Environment for the Service
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// The ARN (Aliyun Resource Name) of the role designated by a user to allow the system to manage his Aliyun resources. If null, use roleArn of role AliyunFCDefaultRole.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The name of the template for the Service
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// Variables for specified template
	TemplateVariables map[string]interface{} `json:"templateVariables,omitempty" xml:"templateVariables,omitempty"`
}

func (s ServiceSpec) String() string {
	return tea.Prettify(s)
}

func (s ServiceSpec) GoString() string {
	return s.String()
}

func (s *ServiceSpec) SetEnvironment(v string) *ServiceSpec {
	s.Environment = &v
	return s
}

func (s *ServiceSpec) SetRoleArn(v string) *ServiceSpec {
	s.RoleArn = &v
	return s
}

func (s *ServiceSpec) SetTemplate(v string) *ServiceSpec {
	s.Template = &v
	return s
}

func (s *ServiceSpec) SetTemplateVariables(v map[string]interface{}) *ServiceSpec {
	s.TemplateVariables = v
	return s
}

type ServiceStatus struct {
	// Credentials required for tasks
	Credentials *StsCredentials `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// A human-readable message indicating details about why the Service is in this condition
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The most recent generation observed
	ObservedGeneration *int32 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	// Time when the last update of the status is observed
	ObservedTime *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	// Details of current state of the Service
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// A simple, high-level summary of where the Service is in its lifecycle
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s ServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s ServiceStatus) GoString() string {
	return s.String()
}

func (s *ServiceStatus) SetCredentials(v *StsCredentials) *ServiceStatus {
	s.Credentials = v
	return s
}

func (s *ServiceStatus) SetMessage(v string) *ServiceStatus {
	s.Message = &v
	return s
}

func (s *ServiceStatus) SetObservedGeneration(v int32) *ServiceStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *ServiceStatus) SetObservedTime(v string) *ServiceStatus {
	s.ObservedTime = &v
	return s
}

func (s *ServiceStatus) SetOutput(v map[string]interface{}) *ServiceStatus {
	s.Output = v
	return s
}

func (s *ServiceStatus) SetPhase(v string) *ServiceStatus {
	s.Phase = &v
	return s
}

type StsCredentials struct {
	// Access key ID
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// Expiration time of the credentials
	ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// The kind of the credentials
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Secret access key
	SecretAccessKey *string `json:"secretAccessKey,omitempty" xml:"secretAccessKey,omitempty"`
	// Token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s StsCredentials) String() string {
	return tea.Prettify(s)
}

func (s StsCredentials) GoString() string {
	return s.String()
}

func (s *StsCredentials) SetAccessKeyId(v string) *StsCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *StsCredentials) SetExpirationTime(v string) *StsCredentials {
	s.ExpirationTime = &v
	return s
}

func (s *StsCredentials) SetKind(v string) *StsCredentials {
	s.Kind = &v
	return s
}

func (s *StsCredentials) SetSecretAccessKey(v string) *StsCredentials {
	s.SecretAccessKey = &v
	return s
}

func (s *StsCredentials) SetToken(v string) *StsCredentials {
	s.Token = &v
	return s
}

type GetEnvironmentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *Environment) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetServiceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetBody(v *Service) *GetServiceResponse {
	s.Body = v
	return s
}

type ListEnvironmentsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*Environment     `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v []*Environment) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type ListServicesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*Service         `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetBody(v []*Service) *ListServicesResponse {
	s.Body = v
	return s
}

type PutEnvironmentRequest struct {
	// An environment for serverless deployments
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *PutEnvironmentRequest) SetBody(v *Environment) *PutEnvironmentRequest {
	s.Body = v
	return s
}

type PutEnvironmentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Environment       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *PutEnvironmentResponse) SetHeaders(v map[string]*string) *PutEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *PutEnvironmentResponse) SetBody(v *Environment) *PutEnvironmentResponse {
	s.Body = v
	return s
}

type PutServiceRequest struct {
	// A service for serverless deployments
	Body *Service `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PutServiceRequest) GoString() string {
	return s.String()
}

func (s *PutServiceRequest) SetBody(v *Service) *PutServiceRequest {
	s.Body = v
	return s
}

type PutServiceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PutServiceResponse) GoString() string {
	return s.String()
}

func (s *PutServiceResponse) SetHeaders(v map[string]*string) *PutServiceResponse {
	s.Headers = v
	return s
}

func (s *PutServiceResponse) SetBody(v *Service) *PutServiceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("serverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironment(name *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(name *string) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironments() (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices() (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutEnvironment(name *string, request *PutEnvironmentRequest) (_result *PutEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutEnvironmentResponse{}
	_body, _err := client.PutEnvironmentWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEnvironmentWithOptions(name *string, request *PutEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnvironment"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/environments/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutService(name *string, request *PutServiceRequest) (_result *PutServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutServiceResponse{}
	_body, _err := client.PutServiceWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutServiceWithOptions(name *string, request *PutServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutService"),
		Version:     tea.String("2021-09-24"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/apis/serverlessdeployment/v1/services/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
