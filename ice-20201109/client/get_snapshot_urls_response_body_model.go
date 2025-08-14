// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotUrlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSnapshotUrlsResponseBody
	GetRequestId() *string
	SetSnapshotUrls(v []*string) *GetSnapshotUrlsResponseBody
	GetSnapshotUrls() []*string
	SetTotal(v int32) *GetSnapshotUrlsResponseBody
	GetTotal() *int32
	SetWebVTTUrl(v string) *GetSnapshotUrlsResponseBody
	GetWebVTTUrl() *string
}

type GetSnapshotUrlsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of snapshot URLs.
	SnapshotUrls []*string `json:"SnapshotUrls,omitempty" xml:"SnapshotUrls,omitempty" type:"Repeated"`
	// The total number of snapshots.
	//
	// example:
	//
	// 30
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The URL of the WebVTT file.
	//
	// example:
	//
	// http://test-bucket.oss-cn-shanghai.aliyuncs.com/ouoput.vtt
	WebVTTUrl *string `json:"WebVTTUrl,omitempty" xml:"WebVTTUrl,omitempty"`
}

func (s GetSnapshotUrlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotUrlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSnapshotUrlsResponseBody) GetSnapshotUrls() []*string {
	return s.SnapshotUrls
}

func (s *GetSnapshotUrlsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetSnapshotUrlsResponseBody) GetWebVTTUrl() *string {
	return s.WebVTTUrl
}

func (s *GetSnapshotUrlsResponseBody) SetRequestId(v string) *GetSnapshotUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSnapshotUrlsResponseBody) SetSnapshotUrls(v []*string) *GetSnapshotUrlsResponseBody {
	s.SnapshotUrls = v
	return s
}

func (s *GetSnapshotUrlsResponseBody) SetTotal(v int32) *GetSnapshotUrlsResponseBody {
	s.Total = &v
	return s
}

func (s *GetSnapshotUrlsResponseBody) SetWebVTTUrl(v string) *GetSnapshotUrlsResponseBody {
	s.WebVTTUrl = &v
	return s
}

func (s *GetSnapshotUrlsResponseBody) Validate() error {
	return dara.Validate(s)
}
