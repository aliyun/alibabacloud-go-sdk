// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentProfilesResponseBody
	GetCode() *string
	SetData(v string) *DeleteAgentProfilesResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteAgentProfilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAgentProfilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentProfilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentProfilesResponseBody
	GetSuccess() *bool
}

type DeleteAgentProfilesResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API prompt message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgentProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentProfilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentProfilesResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAgentProfilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAgentProfilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentProfilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentProfilesResponseBody) SetCode(v string) *DeleteAgentProfilesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) SetData(v string) *DeleteAgentProfilesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) SetHttpStatusCode(v int32) *DeleteAgentProfilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) SetMessage(v string) *DeleteAgentProfilesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) SetRequestId(v string) *DeleteAgentProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) SetSuccess(v bool) *DeleteAgentProfilesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentProfilesResponseBody) Validate() error {
	return dara.Validate(s)
}
