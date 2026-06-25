// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationImageResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationImageResponseBodyData) *DescribeApplicationImageResponseBody
	GetData() *DescribeApplicationImageResponseBodyData
	SetErrorCode(v string) *DescribeApplicationImageResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationImageResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationImageResponseBody
	GetTraceId() *string
}

type DescribeApplicationImageResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client-side error occurred.
	//
	// - **5xx**: A server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application image information.
	Data *DescribeApplicationImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is left empty if the request is successful.
	//
	// - This parameter returns an error code if the request fails. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - **success*	- is returned if the request is successful.
	//
	// - An error code is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationImageResponseBody) GetData() *DescribeApplicationImageResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationImageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationImageResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationImageResponseBody) SetCode(v string) *DescribeApplicationImageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetData(v *DescribeApplicationImageResponseBodyData) *DescribeApplicationImageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetErrorCode(v string) *DescribeApplicationImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetMessage(v string) *DescribeApplicationImageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetRequestId(v string) *DescribeApplicationImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetSuccess(v bool) *DescribeApplicationImageResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetTraceId(v string) *DescribeApplicationImageResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationImageResponseBodyData struct {
	// This parameter is reserved.
	CrUrl *string `json:"CrUrl,omitempty" xml:"CrUrl,omitempty"`
	// This parameter is reserved.
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// demo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The image namespace.
	//
	// example:
	//
	// demo
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The image repository type. Currently, only Container Registry is supported.
	//
	// example:
	//
	// ALI_HUB
	RepoOriginType *string `json:"RepoOriginType,omitempty" xml:"RepoOriginType,omitempty"`
	// The image tag.
	//
	// example:
	//
	// latest
	RepoTag *string `json:"RepoTag,omitempty" xml:"RepoTag,omitempty"`
	// A reserved field.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
}

func (s DescribeApplicationImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponseBodyData) GetCrUrl() *string {
	return s.CrUrl
}

func (s *DescribeApplicationImageResponseBodyData) GetLogo() *string {
	return s.Logo
}

func (s *DescribeApplicationImageResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationImageResponseBodyData) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeApplicationImageResponseBodyData) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeApplicationImageResponseBodyData) GetRepoOriginType() *string {
	return s.RepoOriginType
}

func (s *DescribeApplicationImageResponseBodyData) GetRepoTag() *string {
	return s.RepoTag
}

func (s *DescribeApplicationImageResponseBodyData) GetRepoType() *string {
	return s.RepoType
}

func (s *DescribeApplicationImageResponseBodyData) SetCrUrl(v string) *DescribeApplicationImageResponseBodyData {
	s.CrUrl = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetLogo(v string) *DescribeApplicationImageResponseBodyData {
	s.Logo = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRegionId(v string) *DescribeApplicationImageResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoName(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoName = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoNamespace(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoOriginType(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoOriginType = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoTag(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoTag = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoType(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoType = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
