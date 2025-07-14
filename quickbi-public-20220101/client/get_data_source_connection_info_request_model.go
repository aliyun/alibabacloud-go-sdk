// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceConnectionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDsId(v string) *GetDataSourceConnectionInfoRequest
	GetDsId() *string
}

type GetDataSourceConnectionInfoRequest struct {
	// Data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
}

func (s GetDataSourceConnectionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceConnectionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceConnectionInfoRequest) GetDsId() *string {
	return s.DsId
}

func (s *GetDataSourceConnectionInfoRequest) SetDsId(v string) *GetDataSourceConnectionInfoRequest {
	s.DsId = &v
	return s
}

func (s *GetDataSourceConnectionInfoRequest) Validate() error {
	return dara.Validate(s)
}
