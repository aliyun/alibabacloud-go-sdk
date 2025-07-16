// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchGetFormDataByIdListRequest
	GetAppType() *string
	SetFormInstanceIdList(v []*string) *BatchGetFormDataByIdListRequest
	GetFormInstanceIdList() []*string
	SetFormUuid(v string) *BatchGetFormDataByIdListRequest
	GetFormUuid() *string
	SetNeedFormInstanceValue(v bool) *BatchGetFormDataByIdListRequest
	GetNeedFormInstanceValue() *bool
	SetSystemToken(v string) *BatchGetFormDataByIdListRequest
	GetSystemToken() *string
}

type BatchGetFormDataByIdListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormInstanceIdList []*string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// true
	NeedFormInstanceValue *bool `json:"NeedFormInstanceValue,omitempty" xml:"NeedFormInstanceValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s BatchGetFormDataByIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchGetFormDataByIdListRequest) GetFormInstanceIdList() []*string {
	return s.FormInstanceIdList
}

func (s *BatchGetFormDataByIdListRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchGetFormDataByIdListRequest) GetNeedFormInstanceValue() *bool {
	return s.NeedFormInstanceValue
}

func (s *BatchGetFormDataByIdListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchGetFormDataByIdListRequest) SetAppType(v string) *BatchGetFormDataByIdListRequest {
	s.AppType = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetFormInstanceIdList(v []*string) *BatchGetFormDataByIdListRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetFormUuid(v string) *BatchGetFormDataByIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetNeedFormInstanceValue(v bool) *BatchGetFormDataByIdListRequest {
	s.NeedFormInstanceValue = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetSystemToken(v string) *BatchGetFormDataByIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) Validate() error {
	return dara.Validate(s)
}
