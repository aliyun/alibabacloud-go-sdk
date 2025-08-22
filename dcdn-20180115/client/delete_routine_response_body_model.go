// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DeleteRoutineResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DeleteRoutineResponseBody
	GetRequestId() *string
}

type DeleteRoutineResponseBody struct {
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

func (s DeleteRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DeleteRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineResponseBody) SetRequestId(v string) *DeleteRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}
