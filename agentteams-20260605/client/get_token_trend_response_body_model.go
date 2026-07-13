// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenTrendResponseBody
	GetCode() *string
	SetData(v *GetTokenTrendResponseBodyData) *GetTokenTrendResponseBody
	GetData() *GetTokenTrendResponseBodyData
	SetHttpStatusCode(v int32) *GetTokenTrendResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTokenTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenTrendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTokenTrendResponseBody
	GetSuccess() *bool
}

type GetTokenTrendResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetTokenTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTokenTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenTrendResponseBody) GetData() *GetTokenTrendResponseBodyData {
	return s.Data
}

func (s *GetTokenTrendResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTokenTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenTrendResponseBody) SetCode(v string) *GetTokenTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenTrendResponseBody) SetData(v *GetTokenTrendResponseBodyData) *GetTokenTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenTrendResponseBody) SetHttpStatusCode(v int32) *GetTokenTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTokenTrendResponseBody) SetMessage(v string) *GetTokenTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenTrendResponseBody) SetRequestId(v string) *GetTokenTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenTrendResponseBody) SetSuccess(v bool) *GetTokenTrendResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenTrendResponseBodyData struct {
	GroupBy *string                                `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Series  []*GetTokenTrendResponseBodyDataSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s GetTokenTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenTrendResponseBodyData) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetTokenTrendResponseBodyData) GetSeries() []*GetTokenTrendResponseBodyDataSeries {
	return s.Series
}

func (s *GetTokenTrendResponseBodyData) SetGroupBy(v string) *GetTokenTrendResponseBodyData {
	s.GroupBy = &v
	return s
}

func (s *GetTokenTrendResponseBodyData) SetSeries(v []*GetTokenTrendResponseBodyDataSeries) *GetTokenTrendResponseBodyData {
	s.Series = v
	return s
}

func (s *GetTokenTrendResponseBodyData) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTokenTrendResponseBodyDataSeries struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Name *string       `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTokenTrendResponseBodyDataSeries) String() string {
	return dara.Prettify(s)
}

func (s GetTokenTrendResponseBodyDataSeries) GoString() string {
	return s.String()
}

func (s *GetTokenTrendResponseBodyDataSeries) GetData() []interface{} {
	return s.Data
}

func (s *GetTokenTrendResponseBodyDataSeries) GetName() *string {
	return s.Name
}

func (s *GetTokenTrendResponseBodyDataSeries) SetData(v []interface{}) *GetTokenTrendResponseBodyDataSeries {
	s.Data = v
	return s
}

func (s *GetTokenTrendResponseBodyDataSeries) SetName(v string) *GetTokenTrendResponseBodyDataSeries {
	s.Name = &v
	return s
}

func (s *GetTokenTrendResponseBodyDataSeries) Validate() error {
	return dara.Validate(s)
}
