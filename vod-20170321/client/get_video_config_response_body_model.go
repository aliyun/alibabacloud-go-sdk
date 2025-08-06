// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadSwitch(v string) *GetVideoConfigResponseBody
	GetDownloadSwitch() *string
	SetRequestId(v string) *GetVideoConfigResponseBody
	GetRequestId() *string
}

type GetVideoConfigResponseBody struct {
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoConfigResponseBody) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *GetVideoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoConfigResponseBody) SetDownloadSwitch(v string) *GetVideoConfigResponseBody {
	s.DownloadSwitch = &v
	return s
}

func (s *GetVideoConfigResponseBody) SetRequestId(v string) *GetVideoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
