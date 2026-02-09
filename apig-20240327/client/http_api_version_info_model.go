// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiVersionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *HttpApiVersionInfo
	GetEnable() *bool
	SetHeaderName(v string) *HttpApiVersionInfo
	GetHeaderName() *string
	SetQueryName(v string) *HttpApiVersionInfo
	GetQueryName() *string
	SetScheme(v string) *HttpApiVersionInfo
	GetScheme() *string
	SetVersion(v string) *HttpApiVersionInfo
	GetVersion() *string
}

type HttpApiVersionInfo struct {
	// Specifies whether to enable versioning.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The key in the specified header when the header versioning solution is used.
	//
	// example:
	//
	// my-version
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// The key in the specified query parameter when the query versioning solution is used.
	//
	// example:
	//
	// myVersion
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// The versioning solution.
	//
	// example:
	//
	// Query
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// The version number.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HttpApiVersionInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiVersionInfo) GoString() string {
	return s.String()
}

func (s *HttpApiVersionInfo) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiVersionInfo) GetHeaderName() *string {
	return s.HeaderName
}

func (s *HttpApiVersionInfo) GetQueryName() *string {
	return s.QueryName
}

func (s *HttpApiVersionInfo) GetScheme() *string {
	return s.Scheme
}

func (s *HttpApiVersionInfo) GetVersion() *string {
	return s.Version
}

func (s *HttpApiVersionInfo) SetEnable(v bool) *HttpApiVersionInfo {
	s.Enable = &v
	return s
}

func (s *HttpApiVersionInfo) SetHeaderName(v string) *HttpApiVersionInfo {
	s.HeaderName = &v
	return s
}

func (s *HttpApiVersionInfo) SetQueryName(v string) *HttpApiVersionInfo {
	s.QueryName = &v
	return s
}

func (s *HttpApiVersionInfo) SetScheme(v string) *HttpApiVersionInfo {
	s.Scheme = &v
	return s
}

func (s *HttpApiVersionInfo) SetVersion(v string) *HttpApiVersionInfo {
	s.Version = &v
	return s
}

func (s *HttpApiVersionInfo) Validate() error {
	return dara.Validate(s)
}
