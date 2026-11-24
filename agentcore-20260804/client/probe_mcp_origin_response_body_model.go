// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeMcpOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProbeMcpOriginResponseBody
	GetCode() *string
	SetData(v *ProbeMcpOriginResponseBodyData) *ProbeMcpOriginResponseBody
	GetData() *ProbeMcpOriginResponseBodyData
	SetHttpStatusCode(v int32) *ProbeMcpOriginResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ProbeMcpOriginResponseBody
	GetMessage() *string
	SetRequestId(v string) *ProbeMcpOriginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ProbeMcpOriginResponseBody
	GetSuccess() *bool
}

type ProbeMcpOriginResponseBody struct {
	// The status code of the API operation. SUCCESS indicates that the operation completed normally. For the origin server result, see data.status.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The read-only, one-time, and bounded probe result. It is not persisted and is not associated with subsequently created MCPs.
	Data *ProbeMcpOriginResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the POP API request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The processing result of the API operation, excluding the original response from the origin server.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID used to correlate this API call.
	//
	// example:
	//
	// 8F7A9D02-8E5B-4376-8D2B-6A9CB3A42B16
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API operation was processed successfully. This means that authorization, parameter validation, and execution mechanisms completed normally, which does not equate to a successful origin server check.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ProbeMcpOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBody) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProbeMcpOriginResponseBody) GetData() *ProbeMcpOriginResponseBodyData {
	return s.Data
}

func (s *ProbeMcpOriginResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ProbeMcpOriginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProbeMcpOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProbeMcpOriginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ProbeMcpOriginResponseBody) SetCode(v string) *ProbeMcpOriginResponseBody {
	s.Code = &v
	return s
}

func (s *ProbeMcpOriginResponseBody) SetData(v *ProbeMcpOriginResponseBodyData) *ProbeMcpOriginResponseBody {
	s.Data = v
	return s
}

func (s *ProbeMcpOriginResponseBody) SetHttpStatusCode(v int32) *ProbeMcpOriginResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ProbeMcpOriginResponseBody) SetMessage(v string) *ProbeMcpOriginResponseBody {
	s.Message = &v
	return s
}

func (s *ProbeMcpOriginResponseBody) SetRequestId(v string) *ProbeMcpOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProbeMcpOriginResponseBody) SetSuccess(v bool) *ProbeMcpOriginResponseBody {
	s.Success = &v
	return s
}

func (s *ProbeMcpOriginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProbeMcpOriginResponseBodyData struct {
	// The check coverage. It describes the target, protocol, or HTTP request conditions for this check and does not prove that the formal gateway and other APIs are available.
	Coverage *ProbeMcpOriginResponseBodyDataCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Struct"`
	// Indicates origin server failures. This field is empty upon success. Origin server 401 errors or connection failures are not uniformly mapped to API exceptions.
	Error *ProbeMcpOriginResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The HTTP check result. This is returned only for the HTTP type. The definition and actual request status are separated. This is null for direct connections. A 2xx status code only proves the success of this request and does not prove that business functions or other APIs are available.
	HttpResult *ProbeMcpOriginResponseBodyDataHttpResult `json:"httpResult,omitempty" xml:"httpResult,omitempty" type:"Struct"`
	// The declared and actual test results for controlled endpoints of a single origin server. This is an empty array for the HTTP type.
	Observations []*ProbeMcpOriginResponseBodyDataObservations `json:"observations,omitempty" xml:"observations,omitempty" type:"Repeated"`
	// The ID used only to correlate desensitized logs. It is not used for queries, authorization, binding MCPs, or reusing check evidence.
	//
	// example:
	//
	// mcpprobe-8f7a9d02-8e5b-4376-8d2b-6a9cb3a42b16
	ProbeId *string `json:"probeId,omitempty" xml:"probeId,omitempty"`
	// The observation time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-11-19T08:30:00Z
	ProbedAt *string `json:"probedAt,omitempty" xml:"probedAt,omitempty"`
	// The recommended protocol combination. It only includes verified combinations that the platform can access. This is empty when no combinations are available or for the HTTP type.
	Recommendation *ProbeMcpOriginResponseBodyDataRecommendation `json:"recommendation,omitempty" xml:"recommendation,omitempty" type:"Struct"`
	// The combination actually used for tool queries. Specifying this does not mean verification passed. This is empty for the HTTP type.
	SelectedConfiguration *ProbeMcpOriginResponseBodyDataSelectedConfiguration `json:"selectedConfiguration,omitempty" xml:"selectedConfiguration,omitempty" type:"Struct"`
	// The stage of the check. Valid values: AUTHORIZE, VALIDATE, RATE_LIMIT, NETWORK, NETWORK_CONTEXT, CONNECT, DISCOVER, INITIALIZE, SESSION_CHECK, TOOLS_LIST, HTTP_DEFINITION, HTTP_REQUEST, CLEANUP, EXECUTION, or COMPLETE.
	//
	// example:
	//
	// COMPLETE
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// The check result of the origin server. Valid values: SUCCESS, PARTIAL, or FAILED. SUCCESS indicates that the protocol check passed but no capabilities were declared or no tools capability exists. A completely empty tool list also passes. PARTIAL indicates that the tool query failed or was truncated. For HTTP, PARTIAL indicates that only the definition was checked or required evidence is unknown.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The restricted metadata list for this check. Tools are not executed. An empty array is returned if tools are not declared, the list is completely empty, or the type is HTTP. Refer to toolsStatus and toolsComplete for judgment.
	Tools []*ProbeMcpOriginResponseBodyDataTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// Indicates whether the tool list is complete. The value is true for COMPLETE or EMPTY, and false for TRUNCATED. The value is null if not queried, not declared, or failed. Completeness cannot be determined solely by an empty array.
	ToolsComplete *bool `json:"toolsComplete,omitempty" xml:"toolsComplete,omitempty"`
	// The tool query status. Valid values: NOT_RUN, NOT_DECLARED, EMPTY, COMPLETE, TRUNCATED, or FAILED. If tools are not declared, the value is NOT_DECLARED and tools/list is not executed. For the HTTP type, the value is NOT_RUN.
	//
	// example:
	//
	// COMPLETE
	ToolsStatus *string `json:"toolsStatus,omitempty" xml:"toolsStatus,omitempty"`
	// The coverage and completeness limitations of this check.
	Warnings []*ProbeMcpOriginResponseBodyDataWarnings `json:"warnings,omitempty" xml:"warnings,omitempty" type:"Repeated"`
}

func (s ProbeMcpOriginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyData) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyData) GetCoverage() *ProbeMcpOriginResponseBodyDataCoverage {
	return s.Coverage
}

func (s *ProbeMcpOriginResponseBodyData) GetError() *ProbeMcpOriginResponseBodyDataError {
	return s.Error
}

func (s *ProbeMcpOriginResponseBodyData) GetHttpResult() *ProbeMcpOriginResponseBodyDataHttpResult {
	return s.HttpResult
}

func (s *ProbeMcpOriginResponseBodyData) GetObservations() []*ProbeMcpOriginResponseBodyDataObservations {
	return s.Observations
}

func (s *ProbeMcpOriginResponseBodyData) GetProbeId() *string {
	return s.ProbeId
}

func (s *ProbeMcpOriginResponseBodyData) GetProbedAt() *string {
	return s.ProbedAt
}

func (s *ProbeMcpOriginResponseBodyData) GetRecommendation() *ProbeMcpOriginResponseBodyDataRecommendation {
	return s.Recommendation
}

func (s *ProbeMcpOriginResponseBodyData) GetSelectedConfiguration() *ProbeMcpOriginResponseBodyDataSelectedConfiguration {
	return s.SelectedConfiguration
}

func (s *ProbeMcpOriginResponseBodyData) GetStage() *string {
	return s.Stage
}

func (s *ProbeMcpOriginResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ProbeMcpOriginResponseBodyData) GetTools() []*ProbeMcpOriginResponseBodyDataTools {
	return s.Tools
}

func (s *ProbeMcpOriginResponseBodyData) GetToolsComplete() *bool {
	return s.ToolsComplete
}

func (s *ProbeMcpOriginResponseBodyData) GetToolsStatus() *string {
	return s.ToolsStatus
}

func (s *ProbeMcpOriginResponseBodyData) GetWarnings() []*ProbeMcpOriginResponseBodyDataWarnings {
	return s.Warnings
}

func (s *ProbeMcpOriginResponseBodyData) SetCoverage(v *ProbeMcpOriginResponseBodyDataCoverage) *ProbeMcpOriginResponseBodyData {
	s.Coverage = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetError(v *ProbeMcpOriginResponseBodyDataError) *ProbeMcpOriginResponseBodyData {
	s.Error = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetHttpResult(v *ProbeMcpOriginResponseBodyDataHttpResult) *ProbeMcpOriginResponseBodyData {
	s.HttpResult = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetObservations(v []*ProbeMcpOriginResponseBodyDataObservations) *ProbeMcpOriginResponseBodyData {
	s.Observations = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetProbeId(v string) *ProbeMcpOriginResponseBodyData {
	s.ProbeId = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetProbedAt(v string) *ProbeMcpOriginResponseBodyData {
	s.ProbedAt = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetRecommendation(v *ProbeMcpOriginResponseBodyDataRecommendation) *ProbeMcpOriginResponseBodyData {
	s.Recommendation = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetSelectedConfiguration(v *ProbeMcpOriginResponseBodyDataSelectedConfiguration) *ProbeMcpOriginResponseBodyData {
	s.SelectedConfiguration = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetStage(v string) *ProbeMcpOriginResponseBodyData {
	s.Stage = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetStatus(v string) *ProbeMcpOriginResponseBodyData {
	s.Status = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetTools(v []*ProbeMcpOriginResponseBodyDataTools) *ProbeMcpOriginResponseBodyData {
	s.Tools = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetToolsComplete(v bool) *ProbeMcpOriginResponseBodyData {
	s.ToolsComplete = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetToolsStatus(v string) *ProbeMcpOriginResponseBodyData {
	s.ToolsStatus = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) SetWarnings(v []*ProbeMcpOriginResponseBodyDataWarnings) *ProbeMcpOriginResponseBodyData {
	s.Warnings = v
	return s
}

func (s *ProbeMcpOriginResponseBodyData) Validate() error {
	if s.Coverage != nil {
		if err := s.Coverage.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	if s.HttpResult != nil {
		if err := s.HttpResult.Validate(); err != nil {
			return err
		}
	}
	if s.Observations != nil {
		for _, item := range s.Observations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Recommendation != nil {
		if err := s.Recommendation.Validate(); err != nil {
			return err
		}
	}
	if s.SelectedConfiguration != nil {
		if err := s.SelectedConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Warnings != nil {
		for _, item := range s.Warnings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ProbeMcpOriginResponseBodyDataCoverage struct {
	// The evidence-verified legacy or modern protocol. This field is empty during HTTP interface checks.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// Returns only non-sensitive names and excludes internal platform routing or network identities.
	HeaderNames []*string `json:"headerNames,omitempty" xml:"headerNames,omitempty" type:"Repeated"`
	// The selected HTTP method for this check. This field is empty for direct connections.
	//
	// example:
	//
	// HEAD
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The redacted summary of the selected HTTP path for this check. This field is empty for direct connections.
	//
	// example:
	//
	// /mcp
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The evidence-verified MCP transport. This field is empty during HTTP interface checks.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The summary of the origin server target for this check. It does not contain sensitive query parameters or credentials.
	//
	// example:
	//
	// https://mcp.example.com/mcp
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The origin server check before creation is fixed to false, which does not indicate that the resource is ready.
	ViaGateway *bool `json:"viaGateway,omitempty" xml:"viaGateway,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataCoverage) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataCoverage) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetHeaderNames() []*string {
	return s.HeaderNames
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetMethod() *string {
	return s.Method
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetPath() *string {
	return s.Path
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetProtocol() *string {
	return s.Protocol
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetTarget() *string {
	return s.Target
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) GetViaGateway() *bool {
	return s.ViaGateway
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetBackendMcpProtocol(v string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.BackendMcpProtocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetHeaderNames(v []*string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.HeaderNames = v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetMethod(v string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.Method = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetPath(v string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.Path = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetProtocol(v string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.Protocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetTarget(v string) *ProbeMcpOriginResponseBodyDataCoverage {
	s.Target = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) SetViaGateway(v bool) *ProbeMcpOriginResponseBodyDataCoverage {
	s.ViaGateway = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataCoverage) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataError struct {
	// Returns only bounded diagnostic codes. The original error body from the origin server is not returned.
	//
	// example:
	//
	// PROBE_TIMEOUT
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The trimmed and redacted diagnostic information. This excludes credentials or original responses.
	//
	// example:
	//
	// The probe deadline was exceeded.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the operation is suitable for retry under a restricted policy. This does not trigger automatic retries.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// The error source, which can be MAGIC, ROUTER, NETWORK, ORIGIN, or UNKNOWN. This is determined solely based on actual evidence.
	//
	// example:
	//
	// UNKNOWN
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The detection or publishing stage where the error occurred.
	//
	// example:
	//
	// INITIALIZE
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *ProbeMcpOriginResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *ProbeMcpOriginResponseBodyDataError) GetRetryable() *bool {
	return s.Retryable
}

func (s *ProbeMcpOriginResponseBodyDataError) GetSource() *string {
	return s.Source
}

func (s *ProbeMcpOriginResponseBodyDataError) GetStage() *string {
	return s.Stage
}

func (s *ProbeMcpOriginResponseBodyDataError) SetCode(v string) *ProbeMcpOriginResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataError) SetMessage(v string) *ProbeMcpOriginResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataError) SetRetryable(v bool) *ProbeMcpOriginResponseBodyDataError {
	s.Retryable = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataError) SetSource(v string) *ProbeMcpOriginResponseBodyDataError {
	s.Source = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataError) SetStage(v string) *ProbeMcpOriginResponseBodyDataError {
	s.Stage = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataHttpResult struct {
	// The connection status of the target. Valid values: PASSED, FAILED, NOT_RUN, or UNKNOWN. This is judged only based on the actual target connection evidence for this check.
	//
	// example:
	//
	// PASSED
	ConnectionStatus *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	// The check result of the API definition. Valid values: PASSED, FAILED, or NOT_RUN. The value is NOT_RUN when no OpenAPI specification is available.
	//
	// example:
	//
	// PASSED
	DefinitionStatus *string `json:"definitionStatus,omitempty" xml:"definitionStatus,omitempty"`
	// The actual HTTP status code received. This is empty if the request was not executed or no response was received.
	//
	// example:
	//
	// 405
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataHttpResult) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataHttpResult) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) GetDefinitionStatus() *string {
	return s.DefinitionStatus
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) SetConnectionStatus(v string) *ProbeMcpOriginResponseBodyDataHttpResult {
	s.ConnectionStatus = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) SetDefinitionStatus(v string) *ProbeMcpOriginResponseBodyDataHttpResult {
	s.DefinitionStatus = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) SetHttpStatusCode(v int32) *ProbeMcpOriginResponseBodyDataHttpResult {
	s.HttpStatusCode = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataHttpResult) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataObservations struct {
	// The backend MCP protocol policy. Valid values: legacy or modern.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// Indicates whether the origin server declares support for this combination. This is empty when there is no declaration evidence.
	Declared *bool `json:"declared,omitempty" xml:"declared,omitempty"`
	// The list of declared versions. It retains only the date versions actually declared by the origin server and does not infer the full set from a single negotiation.
	DeclaredVersions []*string `json:"declaredVersions,omitempty" xml:"declaredVersions,omitempty" type:"Repeated"`
	// The desensitized origin server or controlled subsequent endpoint. Sensitive query parameters are not returned.
	//
	// example:
	//
	// https://mcp.example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The negotiated version. For legacy, it is the negotiated version. For modern, it is the selected and verified date version. This is empty when not confirmed.
	//
	// example:
	//
	// 2025-03-26
	NegotiatedVersion *string `json:"negotiatedVersion,omitempty" xml:"negotiatedVersion,omitempty"`
	// The protocol. Valid values: SSE or StreamableHTTP. The modern policy is only combined with StreamableHTTP.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The session mode. Valid values: STATEFUL, STATELESS, NOT_APPLICABLE, or UNKNOWN. A direct legacy connection without a session is not automatically judged as failed.
	//
	// example:
	//
	// STATELESS
	SessionMode *string `json:"sessionMode,omitempty" xml:"sessionMode,omitempty"`
	// The protocol verification status. Valid values: VERIFIED, UNSUPPORTED, FAILED, or UNKNOWN. UNSUPPORTED requires protocol evidence.
	//
	// example:
	//
	// VERIFIED
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataObservations) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataObservations) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetDeclared() *bool {
	return s.Declared
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetDeclaredVersions() []*string {
	return s.DeclaredVersions
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetNegotiatedVersion() *string {
	return s.NegotiatedVersion
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetProtocol() *string {
	return s.Protocol
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetSessionMode() *string {
	return s.SessionMode
}

func (s *ProbeMcpOriginResponseBodyDataObservations) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetBackendMcpProtocol(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.BackendMcpProtocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetDeclared(v bool) *ProbeMcpOriginResponseBodyDataObservations {
	s.Declared = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetDeclaredVersions(v []*string) *ProbeMcpOriginResponseBodyDataObservations {
	s.DeclaredVersions = v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetEndpoint(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.Endpoint = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetNegotiatedVersion(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.NegotiatedVersion = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetProtocol(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.Protocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetSessionMode(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.SessionMode = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) SetVerificationStatus(v string) *ProbeMcpOriginResponseBodyDataObservations {
	s.VerificationStatus = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataObservations) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataRecommendation struct {
	// The backend MCP protocol policy. Valid values: legacy or modern.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// The actual negotiated or selected and verified origin server MCP date version in YYYY-MM-DD format. This is null when not confirmed. It only displays the origin server protocol date metadata. It can be specified during creation and modified during updates. If omitted during an update, the current value is retained. It does not control backend requests or the external protocol of the gateway.
	//
	// example:
	//
	// 2025-03-26
	OriginProtocolVersion *string `json:"originProtocolVersion,omitempty" xml:"originProtocolVersion,omitempty"`
	// The protocol. Valid values: SSE or StreamableHTTP. The modern policy is only combined with StreamableHTTP.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataRecommendation) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataRecommendation) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) GetOriginProtocolVersion() *string {
	return s.OriginProtocolVersion
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) GetProtocol() *string {
	return s.Protocol
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) SetBackendMcpProtocol(v string) *ProbeMcpOriginResponseBodyDataRecommendation {
	s.BackendMcpProtocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) SetOriginProtocolVersion(v string) *ProbeMcpOriginResponseBodyDataRecommendation {
	s.OriginProtocolVersion = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) SetProtocol(v string) *ProbeMcpOriginResponseBodyDataRecommendation {
	s.Protocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataRecommendation) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataSelectedConfiguration struct {
	// The backend MCP protocol policy. Valid values: legacy or modern.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// The actual negotiated or selected and verified origin server MCP date version in YYYY-MM-DD format. This is null when not confirmed. It only displays the origin server protocol date metadata. It can be specified during creation and modified during updates. If omitted during an update, the current value is retained. It does not control backend requests or the external protocol of the gateway.
	//
	// example:
	//
	// 2025-03-26
	OriginProtocolVersion *string `json:"originProtocolVersion,omitempty" xml:"originProtocolVersion,omitempty"`
	// The protocol. Valid values: SSE or StreamableHTTP. The modern policy is only combined with StreamableHTTP.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataSelectedConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataSelectedConfiguration) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) GetOriginProtocolVersion() *string {
	return s.OriginProtocolVersion
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) GetProtocol() *string {
	return s.Protocol
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) SetBackendMcpProtocol(v string) *ProbeMcpOriginResponseBodyDataSelectedConfiguration {
	s.BackendMcpProtocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) SetOriginProtocolVersion(v string) *ProbeMcpOriginResponseBodyDataSelectedConfiguration {
	s.OriginProtocolVersion = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) SetProtocol(v string) *ProbeMcpOriginResponseBodyDataSelectedConfiguration {
	s.Protocol = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataSelectedConfiguration) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataTools struct {
	// The input schema defined by the protocol, represented as a JSON string.
	//
	// example:
	//
	// {"type":"object","properties":{}}
	InputSchema *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// The output schema defined by the protocol, represented as a JSON string.
	//
	// example:
	//
	// {"type":"object","properties":{"total":{"type":"integer"}}}
	OutputSchema *string `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	// The untrusted description text returned by the origin server. It is escaped when displayed and is not executed as an instruction.
	//
	// example:
	//
	// List inventory items
	ToolDescription *string `json:"toolDescription,omitempty" xml:"toolDescription,omitempty"`
	// The declared name of the tool. Only metadata is displayed.
	//
	// example:
	//
	// list_inventory
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataTools) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataTools) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataTools) GetInputSchema() *string {
	return s.InputSchema
}

func (s *ProbeMcpOriginResponseBodyDataTools) GetOutputSchema() *string {
	return s.OutputSchema
}

func (s *ProbeMcpOriginResponseBodyDataTools) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *ProbeMcpOriginResponseBodyDataTools) GetToolName() *string {
	return s.ToolName
}

func (s *ProbeMcpOriginResponseBodyDataTools) SetInputSchema(v string) *ProbeMcpOriginResponseBodyDataTools {
	s.InputSchema = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataTools) SetOutputSchema(v string) *ProbeMcpOriginResponseBodyDataTools {
	s.OutputSchema = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataTools) SetToolDescription(v string) *ProbeMcpOriginResponseBodyDataTools {
	s.ToolDescription = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataTools) SetToolName(v string) *ProbeMcpOriginResponseBodyDataTools {
	s.ToolName = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataTools) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginResponseBodyDataWarnings struct {
	// The structured identifier for the detection limitation.
	//
	// example:
	//
	// TOOLS_TRUNCATED
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Explains limitations such as UNKNOWN, not executed, or truncated. This excludes credentials and original responses.
	//
	// example:
	//
	// The tool preview limit was reached; protocol verification succeeded.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ProbeMcpOriginResponseBodyDataWarnings) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponseBodyDataWarnings) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponseBodyDataWarnings) GetCode() *string {
	return s.Code
}

func (s *ProbeMcpOriginResponseBodyDataWarnings) GetMessage() *string {
	return s.Message
}

func (s *ProbeMcpOriginResponseBodyDataWarnings) SetCode(v string) *ProbeMcpOriginResponseBodyDataWarnings {
	s.Code = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataWarnings) SetMessage(v string) *ProbeMcpOriginResponseBodyDataWarnings {
	s.Message = &v
	return s
}

func (s *ProbeMcpOriginResponseBodyDataWarnings) Validate() error {
	return dara.Validate(s)
}
