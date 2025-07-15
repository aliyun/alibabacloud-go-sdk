// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveSnapshotNotifyConfigResponseBody
	GetDomainName() *string
	SetNotifyAuthKey(v string) *DescribeLiveSnapshotNotifyConfigResponseBody
	GetNotifyAuthKey() *string
	SetNotifyReqAuth(v string) *DescribeLiveSnapshotNotifyConfigResponseBody
	GetNotifyReqAuth() *string
	SetNotifyUrl(v string) *DescribeLiveSnapshotNotifyConfigResponseBody
	GetNotifyUrl() *string
	SetRequestId(v string) *DescribeLiveSnapshotNotifyConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveSnapshotNotifyConfigResponseBody struct {
	// The main streaming domain.
	//
	// example:
	//
	// www.yourdomain***.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback authentication key.
	//
	// example:
	//
	// yourkey
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// Indicates whether callback authentication is enabled. Valid values:
	//
	// 	- **yes**: Callback authentication is enabled.
	//
	// 	- **no**: Callback authentication is disabled.
	//
	// example:
	//
	// yes
	NotifyReqAuth *string `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://callback.yourdomain***.com
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5056369B-D337-499E-B8B7-B761BD37B08A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveSnapshotNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) GetNotifyReqAuth() *string {
	return s.NotifyReqAuth
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) SetDomainName(v string) *DescribeLiveSnapshotNotifyConfigResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) SetNotifyAuthKey(v string) *DescribeLiveSnapshotNotifyConfigResponseBody {
	s.NotifyAuthKey = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) SetNotifyReqAuth(v string) *DescribeLiveSnapshotNotifyConfigResponseBody {
	s.NotifyReqAuth = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) SetNotifyUrl(v string) *DescribeLiveSnapshotNotifyConfigResponseBody {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) SetRequestId(v string) *DescribeLiveSnapshotNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
