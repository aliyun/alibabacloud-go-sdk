// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResponses(v []*BatchResponseBodyResponses) *BatchResponseBody
	GetResponses() []*BatchResponseBodyResponses
}

type BatchResponseBody struct {
	// All responses of the child requests.
	Responses []*BatchResponseBodyResponses `json:"responses,omitempty" xml:"responses,omitempty" type:"Repeated"`
}

func (s BatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchResponseBody) GoString() string {
	return s.String()
}

func (s *BatchResponseBody) GetResponses() []*BatchResponseBodyResponses {
	return s.Responses
}

func (s *BatchResponseBody) SetResponses(v []*BatchResponseBodyResponses) *BatchResponseBody {
	s.Responses = v
	return s
}

func (s *BatchResponseBody) Validate() error {
	if s.Responses != nil {
		for _, item := range s.Responses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchResponseBodyResponses struct {
	// The response parameters of a child request. For more information, see the topic of the corresponding child request.
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the child request. The ID is used to associate a child request with a response.
	//
	// example:
	//
	// 93433894994ad2e1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The returned HTTP status code of a child request. For more information, see the topic of the corresponding child request.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BatchResponseBodyResponses) String() string {
	return dara.Prettify(s)
}

func (s BatchResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *BatchResponseBodyResponses) GetBody() map[string]interface{} {
	return s.Body
}

func (s *BatchResponseBodyResponses) GetId() *string {
	return s.Id
}

func (s *BatchResponseBodyResponses) GetStatus() *int32 {
	return s.Status
}

func (s *BatchResponseBodyResponses) SetBody(v map[string]interface{}) *BatchResponseBodyResponses {
	s.Body = v
	return s
}

func (s *BatchResponseBodyResponses) SetId(v string) *BatchResponseBodyResponses {
	s.Id = &v
	return s
}

func (s *BatchResponseBodyResponses) SetStatus(v int32) *BatchResponseBodyResponses {
	s.Status = &v
	return s
}

func (s *BatchResponseBodyResponses) Validate() error {
	return dara.Validate(s)
}
