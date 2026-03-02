// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucUserEnterpriseListResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*BucEnterprise) *BucUserEnterpriseListResult
	GetData() []*BucEnterprise
	SetRequestId(v string) *BucUserEnterpriseListResult
	GetRequestId() *string
}

type BucUserEnterpriseListResult struct {
	Data      []*BucEnterprise `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BucUserEnterpriseListResult) String() string {
	return dara.Prettify(s)
}

func (s BucUserEnterpriseListResult) GoString() string {
	return s.String()
}

func (s *BucUserEnterpriseListResult) GetData() []*BucEnterprise {
	return s.Data
}

func (s *BucUserEnterpriseListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *BucUserEnterpriseListResult) SetData(v []*BucEnterprise) *BucUserEnterpriseListResult {
	s.Data = v
	return s
}

func (s *BucUserEnterpriseListResult) SetRequestId(v string) *BucUserEnterpriseListResult {
	s.RequestId = &v
	return s
}

func (s *BucUserEnterpriseListResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
