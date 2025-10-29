// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMixStreamListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMixStreamList(v []*DescribeMixStreamListResponseBodyMixStreamList) *DescribeMixStreamListResponseBody
	GetMixStreamList() []*DescribeMixStreamListResponseBodyMixStreamList
	SetRequestId(v string) *DescribeMixStreamListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeMixStreamListResponseBody
	GetTotal() *int32
}

type DescribeMixStreamListResponseBody struct {
	// Details about the stream mixing tasks.
	MixStreamList []*DescribeMixStreamListResponseBodyMixStreamList `json:"MixStreamList,omitempty" xml:"MixStreamList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BC1E78D3-FA8B-4457-DEE2-6093E1232254
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMixStreamListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMixStreamListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListResponseBody) GetMixStreamList() []*DescribeMixStreamListResponseBodyMixStreamList {
	return s.MixStreamList
}

func (s *DescribeMixStreamListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMixStreamListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMixStreamListResponseBody) SetMixStreamList(v []*DescribeMixStreamListResponseBodyMixStreamList) *DescribeMixStreamListResponseBody {
	s.MixStreamList = v
	return s
}

func (s *DescribeMixStreamListResponseBody) SetRequestId(v string) *DescribeMixStreamListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMixStreamListResponseBody) SetTotal(v int32) *DescribeMixStreamListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMixStreamListResponseBody) Validate() error {
	if s.MixStreamList != nil {
		for _, item := range s.MixStreamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMixStreamListResponseBodyMixStreamList struct {
	// The name of the application.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the stream mixing task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-17T08:39:14Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the stream mixing task was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-17T08:39:15Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of input streams.
	//
	// example:
	//
	// 2
	InputStreamNumber *int32 `json:"InputStreamNumber,omitempty" xml:"InputStreamNumber,omitempty"`
	// The ID of the layout.
	//
	// example:
	//
	// USERDEFINED
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The stream mixing template.
	//
	// example:
	//
	// lp_ld
	MixStreamTemplate *string `json:"MixStreamTemplate,omitempty" xml:"MixStreamTemplate,omitempty"`
	// The ID of the stream mixing task. You can specify this parameter in a request to delete the steam mixing task.
	//
	// example:
	//
	// aaf9a50f-c460-3a9b-f180-38dd8f05****
	MixstreamId *string `json:"MixstreamId,omitempty" xml:"MixstreamId,omitempty"`
	// The name of the output stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeMixStreamListResponseBodyMixStreamList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMixStreamListResponseBodyMixStreamList) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetInputStreamNumber() *int32 {
	return s.InputStreamNumber
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetMixStreamTemplate() *string {
	return s.MixStreamTemplate
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetMixstreamId() *string {
	return s.MixstreamId
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetAppName(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.AppName = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetDomainName(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.DomainName = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetGmtCreate(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetGmtModified(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.GmtModified = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetInputStreamNumber(v int32) *DescribeMixStreamListResponseBodyMixStreamList {
	s.InputStreamNumber = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetLayoutId(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.LayoutId = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetMixStreamTemplate(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.MixStreamTemplate = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetMixstreamId(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.MixstreamId = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) SetStreamName(v string) *DescribeMixStreamListResponseBodyMixStreamList {
	s.StreamName = &v
	return s
}

func (s *DescribeMixStreamListResponseBodyMixStreamList) Validate() error {
	return dara.Validate(s)
}
