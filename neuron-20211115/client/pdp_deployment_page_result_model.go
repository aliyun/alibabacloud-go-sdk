// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpDeploymentPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpServiceDeployment) *PdpDeploymentPageResult
	GetData() []*PdpServiceDeployment
	SetRequestId(v string) *PdpDeploymentPageResult
	GetRequestId() *string
	SetTotal(v int64) *PdpDeploymentPageResult
	GetTotal() *int64
}

type PdpDeploymentPageResult struct {
	Data      []*PdpServiceDeployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpDeploymentPageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpDeploymentPageResult) GoString() string {
	return s.String()
}

func (s *PdpDeploymentPageResult) GetData() []*PdpServiceDeployment {
	return s.Data
}

func (s *PdpDeploymentPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpDeploymentPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpDeploymentPageResult) SetData(v []*PdpServiceDeployment) *PdpDeploymentPageResult {
	s.Data = v
	return s
}

func (s *PdpDeploymentPageResult) SetRequestId(v string) *PdpDeploymentPageResult {
	s.RequestId = &v
	return s
}

func (s *PdpDeploymentPageResult) SetTotal(v int64) *PdpDeploymentPageResult {
	s.Total = &v
	return s
}

func (s *PdpDeploymentPageResult) Validate() error {
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
