// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootApRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *RebootApRequest
	GetApMac() *string
	SetAppCode(v string) *RebootApRequest
	GetAppCode() *string
	SetAppName(v string) *RebootApRequest
	GetAppName() *string
}

type RebootApRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s RebootApRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootApRequest) GoString() string {
	return s.String()
}

func (s *RebootApRequest) GetApMac() *string {
	return s.ApMac
}

func (s *RebootApRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *RebootApRequest) GetAppName() *string {
	return s.AppName
}

func (s *RebootApRequest) SetApMac(v string) *RebootApRequest {
	s.ApMac = &v
	return s
}

func (s *RebootApRequest) SetAppCode(v string) *RebootApRequest {
	s.AppCode = &v
	return s
}

func (s *RebootApRequest) SetAppName(v string) *RebootApRequest {
	s.AppName = &v
	return s
}

func (s *RebootApRequest) Validate() error {
	return dara.Validate(s)
}
