// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudDriveServiceMountTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCloudDriveServiceMountTokenResponseBody
	GetRequestId() *string
	SetToken(v *GetCloudDriveServiceMountTokenResponseBodyToken) *GetCloudDriveServiceMountTokenResponseBody
	GetToken() *GetCloudDriveServiceMountTokenResponseBodyToken
}

type GetCloudDriveServiceMountTokenResponseBody struct {
	// example:
	//
	// DC27288A-F9E1-5092-9B5B-71C27D15****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tokens.
	Token *GetCloudDriveServiceMountTokenResponseBodyToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s GetCloudDriveServiceMountTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudDriveServiceMountTokenResponseBody) GetToken() *GetCloudDriveServiceMountTokenResponseBodyToken {
	return s.Token
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetRequestId(v string) *GetCloudDriveServiceMountTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetToken(v *GetCloudDriveServiceMountTokenResponseBodyToken) *GetCloudDriveServiceMountTokenResponseBody {
	s.Token = v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBody) Validate() error {
	if s.Token != nil {
		if err := s.Token.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudDriveServiceMountTokenResponseBodyToken struct {
	// example:
	//
	// h****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// 2022-10-10T04:41:35Z
	ExpiredAfter *string `json:"ExpiredAfter,omitempty" xml:"ExpiredAfter,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The token.
	//
	// example:
	//
	// 7836fa6eced7dc8d54c775k34iu3h4i2kh534f****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The total capacity of the enterprise drive. Unit: GiB
	//
	// example:
	//
	// 6050416754750
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// example:
	//
	// 605089
	UsedSize *int64 `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) String() string {
	return dara.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetDomainId() *string {
	return s.DomainId
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetExpiredAfter() *string {
	return s.ExpiredAfter
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetStatus() *string {
	return s.Status
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetToken() *string {
	return s.Token
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetDomainId(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.DomainId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetExpiredAfter(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.ExpiredAfter = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetStatus(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Status = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetToken(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Token = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetTotalSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.TotalSize = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetUsedSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.UsedSize = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) Validate() error {
	return dara.Validate(s)
}
