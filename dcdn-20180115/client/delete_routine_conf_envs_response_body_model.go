// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineConfEnvsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DeleteRoutineConfEnvsResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DeleteRoutineConfEnvsResponseBody
	GetRequestId() *string
}

type DeleteRoutineConfEnvsResponseBody struct {
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

func (s DeleteRoutineConfEnvsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineConfEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DeleteRoutineConfEnvsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineConfEnvsResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineConfEnvsResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineConfEnvsResponseBody) SetRequestId(v string) *DeleteRoutineConfEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineConfEnvsResponseBody) Validate() error {
	return dara.Validate(s)
}
