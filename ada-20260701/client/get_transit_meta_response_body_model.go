// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *GetTransitMetaResponseBody
	GetDownloadUrl() *string
	SetExpireAt(v int64) *GetTransitMetaResponseBody
	GetExpireAt() *int64
	SetFilePath(v string) *GetTransitMetaResponseBody
	GetFilePath() *string
	SetRequestId(v string) *GetTransitMetaResponseBody
	GetRequestId() *string
	SetSize(v int64) *GetTransitMetaResponseBody
	GetSize() *int64
	SetStatus(v string) *GetTransitMetaResponseBody
	GetStatus() *string
	SetTransitId(v string) *GetTransitMetaResponseBody
	GetTransitId() *string
}

type GetTransitMetaResponseBody struct {
	// The temporary download URL. If Network is not specified, null is returned. If the file is not yet available, the URL may be unavailable. Do not write this URL to logs, store it for a long time, or forward it to unrelated users.
	//
	// example:
	//
	// https://download.example.invalid/code-review.zip?signature=<REDACTED>
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the transit record. The value is a UTC UNIX timestamp in milliseconds, which indicates the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970. You can directly compare it with the current UNIX timestamp in milliseconds without adding or subtracting 8 hours. Do not continue to use the record after this time.
	//
	// example:
	//
	// 1787734800000
	ExpireAt *int64 `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// The opaque object path corresponding to the file. Callers must not parse or construct this value.
	//
	// example:
	//
	// skill-bundle/tenant-demo/user-demo/20260904120000_code-review.zip
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The request ID, which is used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The file size in bytes. If no available file is detected, null may be returned.
	//
	// example:
	//
	// 4096
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the transit file. A value of PENDING indicates that the file is not yet available and the caller can query again later. A value of SUCCESS indicates that the file is available.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The transit ID.
	//
	// example:
	//
	// transit_0123456789abcdef0123456789abcdef
	TransitId *string `json:"TransitId,omitempty" xml:"TransitId,omitempty"`
}

func (s GetTransitMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTransitMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransitMetaResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetTransitMetaResponseBody) GetExpireAt() *int64 {
	return s.ExpireAt
}

func (s *GetTransitMetaResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *GetTransitMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTransitMetaResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *GetTransitMetaResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTransitMetaResponseBody) GetTransitId() *string {
	return s.TransitId
}

func (s *GetTransitMetaResponseBody) SetDownloadUrl(v string) *GetTransitMetaResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetExpireAt(v int64) *GetTransitMetaResponseBody {
	s.ExpireAt = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetFilePath(v string) *GetTransitMetaResponseBody {
	s.FilePath = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetRequestId(v string) *GetTransitMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetSize(v int64) *GetTransitMetaResponseBody {
	s.Size = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetStatus(v string) *GetTransitMetaResponseBody {
	s.Status = &v
	return s
}

func (s *GetTransitMetaResponseBody) SetTransitId(v string) *GetTransitMetaResponseBody {
	s.TransitId = &v
	return s
}

func (s *GetTransitMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
