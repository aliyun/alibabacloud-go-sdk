// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRoutineCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *UploadRoutineCodeResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *UploadRoutineCodeResponseBody
	GetRequestId() *string
}

type UploadRoutineCodeResponseBody struct {
	// The content returned, such as the code version number and information about the code upload.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DFA2027F-86C0-4421-9593-581A7993696C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRoutineCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRoutineCodeResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UploadRoutineCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadRoutineCodeResponseBody) SetContent(v map[string]interface{}) *UploadRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *UploadRoutineCodeResponseBody) SetRequestId(v string) *UploadRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRoutineCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
