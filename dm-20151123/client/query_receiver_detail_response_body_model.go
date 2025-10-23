// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSchema(v string) *QueryReceiverDetailResponseBody
	GetDataSchema() *string
	SetNextStart(v string) *QueryReceiverDetailResponseBody
	GetNextStart() *string
	SetRequestId(v string) *QueryReceiverDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryReceiverDetailResponseBody
	GetTotalCount() *int32
	SetData(v *QueryReceiverDetailResponseBodyData) *QueryReceiverDetailResponseBody
	GetData() *QueryReceiverDetailResponseBodyData
}

type QueryReceiverDetailResponseBody struct {
	// Field name for the Data of recipients
	//
	// example:
	//
	// UserName,NickName,Gender,Birthday,Mobile
	DataSchema *string `json:"DataSchema,omitempty" xml:"DataSchema,omitempty"`
	// Used for pagination. If there are more results, set this returned value to the NextStart in the next request.
	//
	// example:
	//
	// 90f0243616#40test@example.com
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count (deprecated field, kept for historical compatibility)
	//
	// example:
	//
	// 361
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Detailed information
	Data *QueryReceiverDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryReceiverDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBody) GetDataSchema() *string {
	return s.DataSchema
}

func (s *QueryReceiverDetailResponseBody) GetNextStart() *string {
	return s.NextStart
}

func (s *QueryReceiverDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReceiverDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryReceiverDetailResponseBody) GetData() *QueryReceiverDetailResponseBodyData {
	return s.Data
}

func (s *QueryReceiverDetailResponseBody) SetDataSchema(v string) *QueryReceiverDetailResponseBody {
	s.DataSchema = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetNextStart(v string) *QueryReceiverDetailResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetRequestId(v string) *QueryReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetTotalCount(v int32) *QueryReceiverDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetData(v *QueryReceiverDetailResponseBodyData) *QueryReceiverDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryReceiverDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryReceiverDetailResponseBodyData struct {
	Detail []*QueryReceiverDetailResponseBodyDataDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Repeated"`
}

func (s QueryReceiverDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBodyData) GetDetail() []*QueryReceiverDetailResponseBodyDataDetail {
	return s.Detail
}

func (s *QueryReceiverDetailResponseBodyData) SetDetail(v []*QueryReceiverDetailResponseBodyDataDetail) *QueryReceiverDetailResponseBodyData {
	s.Detail = v
	return s
}

func (s *QueryReceiverDetailResponseBodyData) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryReceiverDetailResponseBodyDataDetail struct {
	// Creation Time
	//
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Content
	//
	// example:
	//
	// {\\"Domains\\": [\\"a.example.net\\", \\"b.example.net\\", \\"c.example.net\\", \\"d.example.net\\"]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Recipient address
	//
	// example:
	//
	// a***@example.net
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Creation time in UTC format
	//
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryReceiverDetailResponseBodyDataDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverDetailResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBodyDataDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryReceiverDetailResponseBodyDataDetail) GetData() *string {
	return s.Data
}

func (s *QueryReceiverDetailResponseBodyDataDetail) GetEmail() *string {
	return s.Email
}

func (s *QueryReceiverDetailResponseBodyDataDetail) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetCreateTime(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetData(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Data = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetEmail(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Email = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetUtcCreateTime(v int64) *QueryReceiverDetailResponseBodyDataDetail {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) Validate() error {
	return dara.Validate(s)
}
