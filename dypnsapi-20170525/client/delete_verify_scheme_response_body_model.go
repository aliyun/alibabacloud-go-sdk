// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVerifySchemeResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVerifySchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVerifySchemeResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteVerifySchemeResponseBody
	GetResult() *bool
}

type DeleteVerifySchemeResponseBody struct {
	// The request is successful. For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E3754956-D0B1-5947-962A-AE767D354F01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// 	- **true**: The verification service is deleted.
	//
	// 	- **false**: The verification service failed to be deleted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVerifySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVerifySchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVerifySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVerifySchemeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteVerifySchemeResponseBody) SetCode(v string) *DeleteVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetMessage(v string) *DeleteVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetRequestId(v string) *DeleteVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetResult(v bool) *DeleteVerifySchemeResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) Validate() error {
	return dara.Validate(s)
}
