// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceM3u8JobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitTraceM3u8JobResponseBodyData) *SubmitTraceM3u8JobResponseBody
	GetData() *SubmitTraceM3u8JobResponseBodyData
	SetMessage(v string) *SubmitTraceM3u8JobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTraceM3u8JobResponseBody
	GetRequestId() *string
}

type SubmitTraceM3u8JobResponseBody struct {
	Data *SubmitTraceM3u8JobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTraceM3u8JobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceM3u8JobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTraceM3u8JobResponseBody) GetData() *SubmitTraceM3u8JobResponseBodyData {
	return s.Data
}

func (s *SubmitTraceM3u8JobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTraceM3u8JobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTraceM3u8JobResponseBody) SetData(v *SubmitTraceM3u8JobResponseBodyData) *SubmitTraceM3u8JobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTraceM3u8JobResponseBody) SetMessage(v string) *SubmitTraceM3u8JobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTraceM3u8JobResponseBody) SetRequestId(v string) *SubmitTraceM3u8JobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTraceM3u8JobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitTraceM3u8JobResponseBodyData struct {
	// example:
	//
	// bfb786c639894f4d8064879202****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTraceM3u8JobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceM3u8JobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTraceM3u8JobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTraceM3u8JobResponseBodyData) SetJobId(v string) *SubmitTraceM3u8JobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitTraceM3u8JobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
