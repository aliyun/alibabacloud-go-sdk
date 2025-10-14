// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRebootInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchEventRebootInstanceResponseBody
	GetRequestId() *string
	SetResults(v []*BatchEventRebootInstanceResponseBodyResults) *BatchEventRebootInstanceResponseBody
	GetResults() []*BatchEventRebootInstanceResponseBodyResults
}

type BatchEventRebootInstanceResponseBody struct {
	// Id of the request。
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchEventRebootInstanceResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchEventRebootInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEventRebootInstanceResponseBody) GetResults() []*BatchEventRebootInstanceResponseBodyResults {
	return s.Results
}

func (s *BatchEventRebootInstanceResponseBody) SetRequestId(v string) *BatchEventRebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEventRebootInstanceResponseBody) SetResults(v []*BatchEventRebootInstanceResponseBodyResults) *BatchEventRebootInstanceResponseBody {
	s.Results = v
	return s
}

func (s *BatchEventRebootInstanceResponseBody) Validate() error {
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

type BatchEventRebootInstanceResponseBodyResults struct {
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

func (s BatchEventRebootInstanceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceResponseBodyResults) GetCode() *int32 {
	return s.Code
}

func (s *BatchEventRebootInstanceResponseBodyResults) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventRebootInstanceResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *BatchEventRebootInstanceResponseBodyResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventRebootInstanceResponseBodyResults) SetCode(v int32) *BatchEventRebootInstanceResponseBodyResults {
	s.Code = &v
	return s
}

func (s *BatchEventRebootInstanceResponseBodyResults) SetEventId(v string) *BatchEventRebootInstanceResponseBodyResults {
	s.EventId = &v
	return s
}

func (s *BatchEventRebootInstanceResponseBodyResults) SetMessage(v string) *BatchEventRebootInstanceResponseBodyResults {
	s.Message = &v
	return s
}

func (s *BatchEventRebootInstanceResponseBodyResults) SetResourceId(v string) *BatchEventRebootInstanceResponseBodyResults {
	s.ResourceId = &v
	return s
}

func (s *BatchEventRebootInstanceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
