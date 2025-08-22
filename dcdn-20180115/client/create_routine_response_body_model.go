// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *CreateRoutineResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *CreateRoutineResponseBody
	GetRequestId() *string
}

type CreateRoutineResponseBody struct {
	// The message returned, such as ""Status": "OK"".
	//
	// example:
	//
	// "Status": "OK"
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CreateRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineResponseBody) SetContent(v map[string]interface{}) *CreateRoutineResponseBody {
	s.Content = v
	return s
}

func (s *CreateRoutineResponseBody) SetRequestId(v string) *CreateRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}
