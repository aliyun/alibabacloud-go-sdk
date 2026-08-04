// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunIdByPkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetAliyunIdByPkRequest
	GetAppName() *string
	SetPk(v string) *GetAliyunIdByPkRequest
	GetPk() *string
}

type GetAliyunIdByPkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetAliyunIdByPkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunIdByPkRequest) GoString() string {
	return s.String()
}

func (s *GetAliyunIdByPkRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAliyunIdByPkRequest) GetPk() *string {
	return s.Pk
}

func (s *GetAliyunIdByPkRequest) SetAppName(v string) *GetAliyunIdByPkRequest {
	s.AppName = &v
	return s
}

func (s *GetAliyunIdByPkRequest) SetPk(v string) *GetAliyunIdByPkRequest {
	s.Pk = &v
	return s
}

func (s *GetAliyunIdByPkRequest) Validate() error {
	return dara.Validate(s)
}
