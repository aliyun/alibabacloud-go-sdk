// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeToutiaoLivePublishResponseBodyContent) *DescribeToutiaoLivePublishResponseBody
	GetContent() []*DescribeToutiaoLivePublishResponseBodyContent
	SetDescription(v string) *DescribeToutiaoLivePublishResponseBody
	GetDescription() *string
	SetRequestId(v string) *DescribeToutiaoLivePublishResponseBody
	GetRequestId() *string
}

type DescribeToutiaoLivePublishResponseBody struct {
	// The stream ingest details.
	Content []*DescribeToutiaoLivePublishResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
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

func (s DescribeToutiaoLivePublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePublishResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePublishResponseBody) GetContent() []*DescribeToutiaoLivePublishResponseBodyContent {
	return s.Content
}

func (s *DescribeToutiaoLivePublishResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeToutiaoLivePublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeToutiaoLivePublishResponseBody) SetContent(v []*DescribeToutiaoLivePublishResponseBodyContent) *DescribeToutiaoLivePublishResponseBody {
	s.Content = v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBody) SetDescription(v string) *DescribeToutiaoLivePublishResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBody) SetRequestId(v string) *DescribeToutiaoLivePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBody) Validate() error {
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

type DescribeToutiaoLivePublishResponseBodyContent struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The bitrate. Unit: bit/s.
	//
	// example:
	//
	// 261587
	Bitrate *float32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The bitrate difference.
	//
	// example:
	//
	// 0
	BwDiff *float32 `json:"BwDiff,omitempty" xml:"BwDiff,omitempty"`
	// The name of the content delivery network (CDN) service.
	//
	// example:
	//
	// ali
	CdnName *string `json:"CdnName,omitempty" xml:"CdnName,omitempty"`
	// The ingest domain.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of dropped frames.
	//
	// example:
	//
	// 0
	Flr *float32 `json:"Flr,omitempty" xml:"Flr,omitempty"`
	// The frame rate.
	//
	// example:
	//
	// 74.4
	Fps *float32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The name of the ingested stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1624358970
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeToutiaoLivePublishResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePublishResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetApp() *string {
	return s.App
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetBitrate() *float32 {
	return s.Bitrate
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetBwDiff() *float32 {
	return s.BwDiff
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetCdnName() *string {
	return s.CdnName
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetDomain() *string {
	return s.Domain
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetFlr() *float32 {
	return s.Flr
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetFps() *float32 {
	return s.Fps
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetApp(v string) *DescribeToutiaoLivePublishResponseBodyContent {
	s.App = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetBitrate(v float32) *DescribeToutiaoLivePublishResponseBodyContent {
	s.Bitrate = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetBwDiff(v float32) *DescribeToutiaoLivePublishResponseBodyContent {
	s.BwDiff = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetCdnName(v string) *DescribeToutiaoLivePublishResponseBodyContent {
	s.CdnName = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetDomain(v string) *DescribeToutiaoLivePublishResponseBodyContent {
	s.Domain = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetFlr(v float32) *DescribeToutiaoLivePublishResponseBodyContent {
	s.Flr = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetFps(v float32) *DescribeToutiaoLivePublishResponseBodyContent {
	s.Fps = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetStreamName(v string) *DescribeToutiaoLivePublishResponseBodyContent {
	s.StreamName = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) SetTimestamp(v int64) *DescribeToutiaoLivePublishResponseBodyContent {
	s.Timestamp = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
