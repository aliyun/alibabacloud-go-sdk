// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryOptimizeTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateQueryOptimizeTagResponseBody
	GetCode() *string
	SetData(v bool) *CreateQueryOptimizeTagResponseBody
	GetData() *bool
	SetMessage(v string) *CreateQueryOptimizeTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQueryOptimizeTagResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateQueryOptimizeTagResponseBody
	GetSuccess() *string
}

type CreateQueryOptimizeTagResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the tags were added to the SQL templates.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQueryOptimizeTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryOptimizeTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueryOptimizeTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateQueryOptimizeTagResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateQueryOptimizeTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQueryOptimizeTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQueryOptimizeTagResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateQueryOptimizeTagResponseBody) SetCode(v string) *CreateQueryOptimizeTagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQueryOptimizeTagResponseBody) SetData(v bool) *CreateQueryOptimizeTagResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQueryOptimizeTagResponseBody) SetMessage(v string) *CreateQueryOptimizeTagResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQueryOptimizeTagResponseBody) SetRequestId(v string) *CreateQueryOptimizeTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueryOptimizeTagResponseBody) SetSuccess(v string) *CreateQueryOptimizeTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQueryOptimizeTagResponseBody) Validate() error {
	return dara.Validate(s)
}
