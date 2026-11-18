// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireMs(v int64) *GetTransitMetaRequest
	GetExpireMs() *int64
	SetFilePath(v string) *GetTransitMetaRequest
	GetFilePath() *string
	SetNetwork(v string) *GetTransitMetaRequest
	GetNetwork() *string
	SetTransitId(v string) *GetTransitMetaRequest
	GetTransitId() *string
}

type GetTransitMetaRequest struct {
	// The validity period of the temporary download URL generated for the current request, in milliseconds. The value must be an integer no less than 1000 and is rounded down to the nearest second. If ExpireMs is omitted, the default validity period is 900000 milliseconds, which indicates 15 minutes. A download URL is generated only when Network is specified.
	//
	// example:
	//
	// 900000
	ExpireMs *int64 `json:"ExpireMs,omitempty" xml:"ExpireMs,omitempty"`
	// The opaque object path returned by CreateTransitUploadPolicy. Specify at least one of FilePath and TransitId.
	//
	// example:
	//
	// skill-bundle/tenant-demo/user-demo/20260904120000_code-review.zip
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The network type used for the download URL. Valid values: public and internal. If this parameter is omitted, no download URL is generated.
	//
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The transit ID. Specify at least one of TransitId and FilePath. If both are provided, TransitId takes precedence.
	//
	// example:
	//
	// transit_0123456789abcdef0123456789abcdef
	TransitId *string `json:"TransitId,omitempty" xml:"TransitId,omitempty"`
}

func (s GetTransitMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransitMetaRequest) GoString() string {
	return s.String()
}

func (s *GetTransitMetaRequest) GetExpireMs() *int64 {
	return s.ExpireMs
}

func (s *GetTransitMetaRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetTransitMetaRequest) GetNetwork() *string {
	return s.Network
}

func (s *GetTransitMetaRequest) GetTransitId() *string {
	return s.TransitId
}

func (s *GetTransitMetaRequest) SetExpireMs(v int64) *GetTransitMetaRequest {
	s.ExpireMs = &v
	return s
}

func (s *GetTransitMetaRequest) SetFilePath(v string) *GetTransitMetaRequest {
	s.FilePath = &v
	return s
}

func (s *GetTransitMetaRequest) SetNetwork(v string) *GetTransitMetaRequest {
	s.Network = &v
	return s
}

func (s *GetTransitMetaRequest) SetTransitId(v string) *GetTransitMetaRequest {
	s.TransitId = &v
	return s
}

func (s *GetTransitMetaRequest) Validate() error {
	return dara.Validate(s)
}
