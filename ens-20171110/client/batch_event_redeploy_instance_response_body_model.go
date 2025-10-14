// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRedeployInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchEventRedeployInstanceResponseBody
	GetRequestId() *string
	SetResults(v []*BatchEventRedeployInstanceResponseBodyResults) *BatchEventRedeployInstanceResponseBody
	GetResults() []*BatchEventRedeployInstanceResponseBodyResults
}

type BatchEventRedeployInstanceResponseBody struct {
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchEventRedeployInstanceResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchEventRedeployInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEventRedeployInstanceResponseBody) GetResults() []*BatchEventRedeployInstanceResponseBodyResults {
	return s.Results
}

func (s *BatchEventRedeployInstanceResponseBody) SetRequestId(v string) *BatchEventRedeployInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEventRedeployInstanceResponseBody) SetResults(v []*BatchEventRedeployInstanceResponseBodyResults) *BatchEventRedeployInstanceResponseBody {
	s.Results = v
	return s
}

func (s *BatchEventRedeployInstanceResponseBody) Validate() error {
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

type BatchEventRedeployInstanceResponseBodyResults struct {
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

func (s BatchEventRedeployInstanceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceResponseBodyResults) GetCode() *int32 {
	return s.Code
}

func (s *BatchEventRedeployInstanceResponseBodyResults) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventRedeployInstanceResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *BatchEventRedeployInstanceResponseBodyResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventRedeployInstanceResponseBodyResults) SetCode(v int32) *BatchEventRedeployInstanceResponseBodyResults {
	s.Code = &v
	return s
}

func (s *BatchEventRedeployInstanceResponseBodyResults) SetEventId(v string) *BatchEventRedeployInstanceResponseBodyResults {
	s.EventId = &v
	return s
}

func (s *BatchEventRedeployInstanceResponseBodyResults) SetMessage(v string) *BatchEventRedeployInstanceResponseBodyResults {
	s.Message = &v
	return s
}

func (s *BatchEventRedeployInstanceResponseBodyResults) SetResourceId(v string) *BatchEventRedeployInstanceResponseBodyResults {
	s.ResourceId = &v
	return s
}

func (s *BatchEventRedeployInstanceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
