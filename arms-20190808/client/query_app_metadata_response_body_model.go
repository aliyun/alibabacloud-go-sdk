// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryAppMetadataResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *QueryAppMetadataResponseBody
	GetData() map[string]interface{}
	SetHttpStatusCode(v int32) *QueryAppMetadataResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryAppMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAppMetadataResponseBody
	GetSuccess() *bool
}

type QueryAppMetadataResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// 	- 2XX: The request is successful.
	//
	// 	- 3XX: A redirection message is returned.
	//
	// 	- 4XX: The request is invalid.
	//
	// 	- 5XX: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	//
	// example:
	//
	// {\\"408d5533\\": \\"SELECT 	- FROM user_base_info\\"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 51877BAC-330C-5845-BDFD-C7859AD33FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAppMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppMetadataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryAppMetadataResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryAppMetadataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryAppMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAppMetadataResponseBody) SetCode(v int32) *QueryAppMetadataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppMetadataResponseBody) SetData(v map[string]interface{}) *QueryAppMetadataResponseBody {
	s.Data = v
	return s
}

func (s *QueryAppMetadataResponseBody) SetHttpStatusCode(v int32) *QueryAppMetadataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAppMetadataResponseBody) SetRequestId(v string) *QueryAppMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppMetadataResponseBody) SetSuccess(v bool) *QueryAppMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAppMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
