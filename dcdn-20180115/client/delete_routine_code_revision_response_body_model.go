// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeRevisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DeleteRoutineCodeRevisionResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DeleteRoutineCodeRevisionResponseBody
	GetRequestId() *string
}

type DeleteRoutineCodeRevisionResponseBody struct {
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

func (s DeleteRoutineCodeRevisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DeleteRoutineCodeRevisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineCodeRevisionResponseBody) SetRequestId(v string) *DeleteRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineCodeRevisionResponseBody) Validate() error {
	return dara.Validate(s)
}
