// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalitySettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeLocalitySettingResponseBody
	GetCode() *int32
	SetData(v *DescribeLocalitySettingResponseBodyData) *DescribeLocalitySettingResponseBody
	GetData() *DescribeLocalitySettingResponseBodyData
	SetHttpStatusCode(v int32) *DescribeLocalitySettingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeLocalitySettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeLocalitySettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLocalitySettingResponseBody
	GetSuccess() *bool
}

type DescribeLocalitySettingResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeLocalitySettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1053-08e4-47a5-b2ab-5c0323de****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLocalitySettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalitySettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLocalitySettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLocalitySettingResponseBody) GetData() *DescribeLocalitySettingResponseBodyData {
	return s.Data
}

func (s *DescribeLocalitySettingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeLocalitySettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLocalitySettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLocalitySettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLocalitySettingResponseBody) SetCode(v int32) *DescribeLocalitySettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLocalitySettingResponseBody) SetData(v *DescribeLocalitySettingResponseBodyData) *DescribeLocalitySettingResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLocalitySettingResponseBody) SetHttpStatusCode(v int32) *DescribeLocalitySettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeLocalitySettingResponseBody) SetMessage(v string) *DescribeLocalitySettingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLocalitySettingResponseBody) SetRequestId(v string) *DescribeLocalitySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLocalitySettingResponseBody) SetSuccess(v bool) *DescribeLocalitySettingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLocalitySettingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLocalitySettingResponseBodyData struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 15
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLocalitySettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalitySettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLocalitySettingResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeLocalitySettingResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DescribeLocalitySettingResponseBodyData) SetEnabled(v bool) *DescribeLocalitySettingResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *DescribeLocalitySettingResponseBodyData) SetThreshold(v float32) *DescribeLocalitySettingResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *DescribeLocalitySettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
