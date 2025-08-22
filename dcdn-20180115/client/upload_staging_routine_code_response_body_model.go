// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStagingRoutineCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *UploadStagingRoutineCodeResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *UploadStagingRoutineCodeResponseBody
	GetRequestId() *string
}

type UploadStagingRoutineCodeResponseBody struct {
	// The parameters required by the code.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DFA2027F-86C0-4421-9593-581A7993696C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadStagingRoutineCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadStagingRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UploadStagingRoutineCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadStagingRoutineCodeResponseBody) SetContent(v map[string]interface{}) *UploadStagingRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *UploadStagingRoutineCodeResponseBody) SetRequestId(v string) *UploadStagingRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadStagingRoutineCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
