// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeRevisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *PublishRoutineCodeRevisionResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *PublishRoutineCodeRevisionResponseBody
	GetRequestId() *string
}

type PublishRoutineCodeRevisionResponseBody struct {
	// The version of the routine code that is published to the specified environment.
	//
	// example:
	//
	// "CodeRevision": "1620876959997924701"
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A513734D-D17B-411E-864D-XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishRoutineCodeRevisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *PublishRoutineCodeRevisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *PublishRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *PublishRoutineCodeRevisionResponseBody) SetRequestId(v string) *PublishRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRoutineCodeRevisionResponseBody) Validate() error {
	return dara.Validate(s)
}
