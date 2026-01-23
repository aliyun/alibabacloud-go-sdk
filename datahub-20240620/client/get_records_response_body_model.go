// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*GetRecordsResponseBodyRecords) *GetRecordsResponseBody
	GetRecords() []*GetRecordsResponseBodyRecords
	SetRequestId(v string) *GetRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRecordsResponseBody
	GetSuccess() *bool
}

type GetRecordsResponseBody struct {
	Records []*GetRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordsResponseBody) GetRecords() []*GetRecordsResponseBodyRecords {
	return s.Records
}

func (s *GetRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRecordsResponseBody) SetRecords(v []*GetRecordsResponseBodyRecords) *GetRecordsResponseBody {
	s.Records = v
	return s
}

func (s *GetRecordsResponseBody) SetRequestId(v string) *GetRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordsResponseBody) SetSuccess(v bool) *GetRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *GetRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRecordsResponseBodyRecords struct {
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Data       []*string          `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1769065251123
	SystemTime *int64 `json:"SystemTime,omitempty" xml:"SystemTime,omitempty"`
}

func (s GetRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s GetRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *GetRecordsResponseBodyRecords) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *GetRecordsResponseBodyRecords) GetData() []*string {
	return s.Data
}

func (s *GetRecordsResponseBodyRecords) GetSystemTime() *int64 {
	return s.SystemTime
}

func (s *GetRecordsResponseBodyRecords) SetAttributes(v map[string]*string) *GetRecordsResponseBodyRecords {
	s.Attributes = v
	return s
}

func (s *GetRecordsResponseBodyRecords) SetData(v []*string) *GetRecordsResponseBodyRecords {
	s.Data = v
	return s
}

func (s *GetRecordsResponseBodyRecords) SetSystemTime(v int64) *GetRecordsResponseBodyRecords {
	s.SystemTime = &v
	return s
}

func (s *GetRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
