// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePlayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeToutiaoLivePlayResponseBodyContent) *DescribeToutiaoLivePlayResponseBody
	GetContent() []*DescribeToutiaoLivePlayResponseBodyContent
	SetDescription(v string) *DescribeToutiaoLivePlayResponseBody
	GetDescription() *string
	SetRequestId(v string) *DescribeToutiaoLivePlayResponseBody
	GetRequestId() *string
}

type DescribeToutiaoLivePlayResponseBody struct {
	// The information about the live stream.
	Content []*DescribeToutiaoLivePlayResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The description of the response status.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeToutiaoLivePlayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePlayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePlayResponseBody) GetContent() []*DescribeToutiaoLivePlayResponseBodyContent {
	return s.Content
}

func (s *DescribeToutiaoLivePlayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeToutiaoLivePlayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeToutiaoLivePlayResponseBody) SetContent(v []*DescribeToutiaoLivePlayResponseBodyContent) *DescribeToutiaoLivePlayResponseBody {
	s.Content = v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBody) SetDescription(v string) *DescribeToutiaoLivePlayResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBody) SetRequestId(v string) *DescribeToutiaoLivePlayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeToutiaoLivePlayResponseBodyContent struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 0.0801239013671875
	Bandwidth *float32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Content Delivery Network (CDN) name.
	//
	// example:
	//
	// ali
	CdnName *string `json:"CdnName,omitempty" xml:"CdnName,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of viewers.
	//
	// example:
	//
	// 452
	PlayNum *int64 `json:"PlayNum,omitempty" xml:"PlayNum,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1625484600
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeToutiaoLivePlayResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePlayResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetApp() *string {
	return s.App
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetBandwidth() *float32 {
	return s.Bandwidth
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetCdnName() *string {
	return s.CdnName
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetDomain() *string {
	return s.Domain
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetPlayNum() *int64 {
	return s.PlayNum
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetApp(v string) *DescribeToutiaoLivePlayResponseBodyContent {
	s.App = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetBandwidth(v float32) *DescribeToutiaoLivePlayResponseBodyContent {
	s.Bandwidth = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetCdnName(v string) *DescribeToutiaoLivePlayResponseBodyContent {
	s.CdnName = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetDomain(v string) *DescribeToutiaoLivePlayResponseBodyContent {
	s.Domain = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetPlayNum(v int64) *DescribeToutiaoLivePlayResponseBodyContent {
	s.PlayNum = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetStreamName(v string) *DescribeToutiaoLivePlayResponseBodyContent {
	s.StreamName = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) SetTimestamp(v int64) *DescribeToutiaoLivePlayResponseBodyContent {
	s.Timestamp = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
