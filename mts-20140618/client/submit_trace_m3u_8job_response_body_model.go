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
	// DEB915C5-D001-5C17-AF65-FF8A65DFE432
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTraceM3u8JobResponseBodyData struct {
	// example:
	//
	// bfb786c639894f4d80648792021e****
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
