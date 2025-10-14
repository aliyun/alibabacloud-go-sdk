// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventMigrateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchEventMigrateInstanceResponseBody
	GetRequestId() *string
	SetResults(v []*BatchEventMigrateInstanceResponseBodyResults) *BatchEventMigrateInstanceResponseBody
	GetResults() []*BatchEventMigrateInstanceResponseBodyResults
}

type BatchEventMigrateInstanceResponseBody struct {
	// example:
	//
	// FF53E96D-3F1A-42F0-8373-1C2B39D72D44
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchEventMigrateInstanceResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchEventMigrateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEventMigrateInstanceResponseBody) GetResults() []*BatchEventMigrateInstanceResponseBodyResults {
	return s.Results
}

func (s *BatchEventMigrateInstanceResponseBody) SetRequestId(v string) *BatchEventMigrateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEventMigrateInstanceResponseBody) SetResults(v []*BatchEventMigrateInstanceResponseBodyResults) *BatchEventMigrateInstanceResponseBody {
	s.Results = v
	return s
}

func (s *BatchEventMigrateInstanceResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchEventMigrateInstanceResponseBodyResults struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// e-d71ff150945b9c02eb6ebc0016328468
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// i-55qi8m11rr53c4i964md8a00l
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s BatchEventMigrateInstanceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceResponseBodyResults) GetCode() *int32 {
	return s.Code
}

func (s *BatchEventMigrateInstanceResponseBodyResults) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventMigrateInstanceResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *BatchEventMigrateInstanceResponseBodyResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventMigrateInstanceResponseBodyResults) SetCode(v int32) *BatchEventMigrateInstanceResponseBodyResults {
	s.Code = &v
	return s
}

func (s *BatchEventMigrateInstanceResponseBodyResults) SetEventId(v string) *BatchEventMigrateInstanceResponseBodyResults {
	s.EventId = &v
	return s
}

func (s *BatchEventMigrateInstanceResponseBodyResults) SetMessage(v string) *BatchEventMigrateInstanceResponseBodyResults {
	s.Message = &v
	return s
}

func (s *BatchEventMigrateInstanceResponseBodyResults) SetResourceId(v string) *BatchEventMigrateInstanceResponseBodyResults {
	s.ResourceId = &v
	return s
}

func (s *BatchEventMigrateInstanceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
