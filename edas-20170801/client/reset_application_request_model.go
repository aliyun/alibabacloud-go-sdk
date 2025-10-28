// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ResetApplicationRequest
	GetAppId() *string
	SetEccInfo(v string) *ResetApplicationRequest
	GetEccInfo() *string
}

type ResetApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92*********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute container (ECC) for which you want to reset the application. Separate multiple ECC IDs with commas (,). You can call the QueryApplicationStatus operation to query the ECC ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0cf49a6c-***********
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s ResetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationRequest) GoString() string {
	return s.String()
}

func (s *ResetApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ResetApplicationRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *ResetApplicationRequest) SetAppId(v string) *ResetApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ResetApplicationRequest) SetEccInfo(v string) *ResetApplicationRequest {
	s.EccInfo = &v
	return s
}

func (s *ResetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
