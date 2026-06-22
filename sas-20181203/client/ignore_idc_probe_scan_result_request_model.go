// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreIdcProbeScanResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreAction(v int32) *IgnoreIdcProbeScanResultRequest
	GetIgnoreAction() *int32
	SetScanResultIds(v string) *IgnoreIdcProbeScanResultRequest
	GetScanResultIds() *string
}

type IgnoreIdcProbeScanResultRequest struct {
	// The action to perform. Valid values:
	//
	// - **1**: whitelist
	//
	// - **2**: ignore.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IgnoreAction *int32 `json:"IgnoreAction,omitempty" xml:"IgnoreAction,omitempty"`
	// The IDs of scan results. Separate multiple IDs with commas (,).
	//
	// > Call the [DescribeIdcProbeScanResultList](~~DescribeIdcProbeScanResultList~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 332098932,332098964,332098963
	ScanResultIds *string `json:"ScanResultIds,omitempty" xml:"ScanResultIds,omitempty"`
}

func (s IgnoreIdcProbeScanResultRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreIdcProbeScanResultRequest) GoString() string {
	return s.String()
}

func (s *IgnoreIdcProbeScanResultRequest) GetIgnoreAction() *int32 {
	return s.IgnoreAction
}

func (s *IgnoreIdcProbeScanResultRequest) GetScanResultIds() *string {
	return s.ScanResultIds
}

func (s *IgnoreIdcProbeScanResultRequest) SetIgnoreAction(v int32) *IgnoreIdcProbeScanResultRequest {
	s.IgnoreAction = &v
	return s
}

func (s *IgnoreIdcProbeScanResultRequest) SetScanResultIds(v string) *IgnoreIdcProbeScanResultRequest {
	s.ScanResultIds = &v
	return s
}

func (s *IgnoreIdcProbeScanResultRequest) Validate() error {
	return dara.Validate(s)
}
