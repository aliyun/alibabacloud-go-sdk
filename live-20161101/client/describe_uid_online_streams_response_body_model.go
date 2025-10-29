// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUidOnlineStreamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeUidOnlineStreamsResponseBodyData) *DescribeUidOnlineStreamsResponseBody
	GetData() []*DescribeUidOnlineStreamsResponseBodyData
	SetRequestId(v string) *DescribeUidOnlineStreamsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeUidOnlineStreamsResponseBody
	GetTotalNum() *int64
}

type DescribeUidOnlineStreamsResponseBody struct {
	Data      []*DescribeUidOnlineStreamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int64                                      `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribeUidOnlineStreamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUidOnlineStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUidOnlineStreamsResponseBody) GetData() []*DescribeUidOnlineStreamsResponseBodyData {
	return s.Data
}

func (s *DescribeUidOnlineStreamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUidOnlineStreamsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeUidOnlineStreamsResponseBody) SetData(v []*DescribeUidOnlineStreamsResponseBodyData) *DescribeUidOnlineStreamsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBody) SetRequestId(v string) *DescribeUidOnlineStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBody) SetTotalNum(v int64) *DescribeUidOnlineStreamsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUidOnlineStreamsResponseBodyData struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeUidOnlineStreamsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUidOnlineStreamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUidOnlineStreamsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeUidOnlineStreamsResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUidOnlineStreamsResponseBodyData) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeUidOnlineStreamsResponseBodyData) SetAppName(v string) *DescribeUidOnlineStreamsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBodyData) SetDomainName(v string) *DescribeUidOnlineStreamsResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBodyData) SetStreamName(v string) *DescribeUidOnlineStreamsResponseBodyData {
	s.StreamName = &v
	return s
}

func (s *DescribeUidOnlineStreamsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
