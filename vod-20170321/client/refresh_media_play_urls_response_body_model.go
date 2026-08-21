// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshMediaPlayUrlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenMediaIds(v string) *RefreshMediaPlayUrlsResponseBody
	GetForbiddenMediaIds() *string
	SetMediaRefreshJobId(v string) *RefreshMediaPlayUrlsResponseBody
	GetMediaRefreshJobId() *string
	SetNonExistMediaIds(v string) *RefreshMediaPlayUrlsResponseBody
	GetNonExistMediaIds() *string
	SetRequestId(v string) *RefreshMediaPlayUrlsResponseBody
	GetRequestId() *string
}

type RefreshMediaPlayUrlsResponseBody struct {
	// The list of audio or video IDs that are forbidden. This is typically because you do not have multi-application permissions. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// a6e49sfgd23p5g9ja7095863****
	ForbiddenMediaIds *string `json:"ForbiddenMediaIds,omitempty" xml:"ForbiddenMediaIds,omitempty"`
	// The ID of the refresh or prefetch task.
	//
	// example:
	//
	// 41d465e31957****
	MediaRefreshJobId *string `json:"MediaRefreshJobId,omitempty" xml:"MediaRefreshJobId,omitempty"`
	// The list of audio or video IDs that do not exist.
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****
	NonExistMediaIds *string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-04D5-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshMediaPlayUrlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshMediaPlayUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshMediaPlayUrlsResponseBody) GetForbiddenMediaIds() *string {
	return s.ForbiddenMediaIds
}

func (s *RefreshMediaPlayUrlsResponseBody) GetMediaRefreshJobId() *string {
	return s.MediaRefreshJobId
}

func (s *RefreshMediaPlayUrlsResponseBody) GetNonExistMediaIds() *string {
	return s.NonExistMediaIds
}

func (s *RefreshMediaPlayUrlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshMediaPlayUrlsResponseBody) SetForbiddenMediaIds(v string) *RefreshMediaPlayUrlsResponseBody {
	s.ForbiddenMediaIds = &v
	return s
}

func (s *RefreshMediaPlayUrlsResponseBody) SetMediaRefreshJobId(v string) *RefreshMediaPlayUrlsResponseBody {
	s.MediaRefreshJobId = &v
	return s
}

func (s *RefreshMediaPlayUrlsResponseBody) SetNonExistMediaIds(v string) *RefreshMediaPlayUrlsResponseBody {
	s.NonExistMediaIds = &v
	return s
}

func (s *RefreshMediaPlayUrlsResponseBody) SetRequestId(v string) *RefreshMediaPlayUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshMediaPlayUrlsResponseBody) Validate() error {
	return dara.Validate(s)
}
