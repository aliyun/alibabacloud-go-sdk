// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDataAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SyncDataAssetsRequest
	GetLang() *string
	SetProductCode(v string) *SyncDataAssetsRequest
	GetProductCode() *string
}

type SyncDataAssetsRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s SyncDataAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *SyncDataAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *SyncDataAssetsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *SyncDataAssetsRequest) SetLang(v string) *SyncDataAssetsRequest {
	s.Lang = &v
	return s
}

func (s *SyncDataAssetsRequest) SetProductCode(v string) *SyncDataAssetsRequest {
	s.ProductCode = &v
	return s
}

func (s *SyncDataAssetsRequest) Validate() error {
	return dara.Validate(s)
}
