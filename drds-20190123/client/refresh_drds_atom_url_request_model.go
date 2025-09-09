// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDrdsAtomUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *RefreshDrdsAtomUrlRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *RefreshDrdsAtomUrlRequest
	GetDrdsInstanceId() *string
}

type RefreshDrdsAtomUrlRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RefreshDrdsAtomUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDrdsAtomUrlRequest) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlRequest) GetDbName() *string {
	return s.DbName
}

func (s *RefreshDrdsAtomUrlRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RefreshDrdsAtomUrlRequest) SetDbName(v string) *RefreshDrdsAtomUrlRequest {
	s.DbName = &v
	return s
}

func (s *RefreshDrdsAtomUrlRequest) SetDrdsInstanceId(v string) *RefreshDrdsAtomUrlRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RefreshDrdsAtomUrlRequest) Validate() error {
	return dara.Validate(s)
}
