// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MultiModalAgentSSEResponseBody
	GetCode() *string
	SetData(v *MultiModalAgentSSEResponseBodyData) *MultiModalAgentSSEResponseBody
	GetData() *MultiModalAgentSSEResponseBodyData
	SetMessage(v string) *MultiModalAgentSSEResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalAgentSSEResponseBody
	GetRequestId() *string
}

type MultiModalAgentSSEResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *MultiModalAgentSSEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalAgentSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentSSEResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalAgentSSEResponseBody) GetCode() *string {
	return s.Code
}

func (s *MultiModalAgentSSEResponseBody) GetData() *MultiModalAgentSSEResponseBodyData {
	return s.Data
}

func (s *MultiModalAgentSSEResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalAgentSSEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalAgentSSEResponseBody) SetCode(v string) *MultiModalAgentSSEResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalAgentSSEResponseBody) SetData(v *MultiModalAgentSSEResponseBodyData) *MultiModalAgentSSEResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalAgentSSEResponseBody) SetMessage(v string) *MultiModalAgentSSEResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalAgentSSEResponseBody) SetRequestId(v string) *MultiModalAgentSSEResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalAgentSSEResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalAgentSSEResponseBodyData struct {
	// The timestamp when the session was created.
	//
	// example:
	//
	// 1783328536
	Created *int64 `json:"Created,omitempty" xml:"Created,omitempty"`
	// The value of dataId passed in the API request. This field is not returned if dataId is not specified in the request.
	//
	// example:
	//
	// dataId-XXX
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// If streaming output is used, this field is null during generation. When generation ends, this field is set to stop if the generation stopped due to a stop token.
	//
	// example:
	//
	// stop
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// The output result.
	//
	// example:
	//
	// "违规，原因XX"
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The credits usage.
	Usage *MultiModalAgentSSEResponseBodyDataUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s MultiModalAgentSSEResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentSSEResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalAgentSSEResponseBodyData) GetCreated() *int64 {
	return s.Created
}

func (s *MultiModalAgentSSEResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalAgentSSEResponseBodyData) GetFinishReason() *string {
	return s.FinishReason
}

func (s *MultiModalAgentSSEResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *MultiModalAgentSSEResponseBodyData) GetUsage() *MultiModalAgentSSEResponseBodyDataUsage {
	return s.Usage
}

func (s *MultiModalAgentSSEResponseBodyData) SetCreated(v int64) *MultiModalAgentSSEResponseBodyData {
	s.Created = &v
	return s
}

func (s *MultiModalAgentSSEResponseBodyData) SetDataId(v string) *MultiModalAgentSSEResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalAgentSSEResponseBodyData) SetFinishReason(v string) *MultiModalAgentSSEResponseBodyData {
	s.FinishReason = &v
	return s
}

func (s *MultiModalAgentSSEResponseBodyData) SetOutput(v string) *MultiModalAgentSSEResponseBodyData {
	s.Output = &v
	return s
}

func (s *MultiModalAgentSSEResponseBodyData) SetUsage(v *MultiModalAgentSSEResponseBodyDataUsage) *MultiModalAgentSSEResponseBodyData {
	s.Usage = v
	return s
}

func (s *MultiModalAgentSSEResponseBodyData) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalAgentSSEResponseBodyDataUsage struct {
	// The number of credits consumed.
	//
	// example:
	//
	// 1.23
	Credits *float64 `json:"Credits,omitempty" xml:"Credits,omitempty"`
}

func (s MultiModalAgentSSEResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentSSEResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *MultiModalAgentSSEResponseBodyDataUsage) GetCredits() *float64 {
	return s.Credits
}

func (s *MultiModalAgentSSEResponseBodyDataUsage) SetCredits(v float64) *MultiModalAgentSSEResponseBodyDataUsage {
	s.Credits = &v
	return s
}

func (s *MultiModalAgentSSEResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
