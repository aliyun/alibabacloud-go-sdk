// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstantScoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstantScoreResponseBody
	GetCode() *string
	SetData(v *GetInstantScoreResponseBodyData) *GetInstantScoreResponseBody
	GetData() *GetInstantScoreResponseBodyData
	SetMessage(v string) *GetInstantScoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstantScoreResponseBody
	GetRequestId() *string
}

type GetInstantScoreResponseBody struct {
	// 集群ID
	//
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetInstantScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Query no data
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInstantScoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstantScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstantScoreResponseBody) GetData() *GetInstantScoreResponseBodyData {
	return s.Data
}

func (s *GetInstantScoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstantScoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstantScoreResponseBody) SetCode(v string) *GetInstantScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstantScoreResponseBody) SetData(v *GetInstantScoreResponseBodyData) *GetInstantScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetInstantScoreResponseBody) SetMessage(v string) *GetInstantScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstantScoreResponseBody) SetRequestId(v string) *GetInstantScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstantScoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstantScoreResponseBodyData struct {
	// example:
	//
	// 100
	Error *float32 `json:"error,omitempty" xml:"error,omitempty"`
	// example:
	//
	// 100
	Latency *float32 `json:"latency,omitempty" xml:"latency,omitempty"`
	// example:
	//
	// 100
	Load *float32 `json:"load,omitempty" xml:"load,omitempty"`
	// example:
	//
	// 100
	Saturation *float32 `json:"saturation,omitempty" xml:"saturation,omitempty"`
	Total      *float32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetInstantScoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstantScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponseBodyData) GetError() *float32 {
	return s.Error
}

func (s *GetInstantScoreResponseBodyData) GetLatency() *float32 {
	return s.Latency
}

func (s *GetInstantScoreResponseBodyData) GetLoad() *float32 {
	return s.Load
}

func (s *GetInstantScoreResponseBodyData) GetSaturation() *float32 {
	return s.Saturation
}

func (s *GetInstantScoreResponseBodyData) GetTotal() *float32 {
	return s.Total
}

func (s *GetInstantScoreResponseBodyData) SetError(v float32) *GetInstantScoreResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetLatency(v float32) *GetInstantScoreResponseBodyData {
	s.Latency = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetLoad(v float32) *GetInstantScoreResponseBodyData {
	s.Load = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetSaturation(v float32) *GetInstantScoreResponseBodyData {
	s.Saturation = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetTotal(v float32) *GetInstantScoreResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
