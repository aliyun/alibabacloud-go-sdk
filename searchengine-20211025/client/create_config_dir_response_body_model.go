// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDirResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateConfigDirResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateConfigDirResponseBody
	GetResult() map[string]interface{}
}

type CreateConfigDirResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateConfigDirResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigDirResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigDirResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateConfigDirResponseBody) SetRequestId(v string) *CreateConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigDirResponseBody) SetResult(v map[string]interface{}) *CreateConfigDirResponseBody {
	s.Result = v
	return s
}

func (s *CreateConfigDirResponseBody) Validate() error {
	return dara.Validate(s)
}
