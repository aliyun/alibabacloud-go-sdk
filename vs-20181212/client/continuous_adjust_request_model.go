// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousAdjustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFocus(v string) *ContinuousAdjustRequest
	GetFocus() *string
	SetId(v string) *ContinuousAdjustRequest
	GetId() *string
	SetIris(v string) *ContinuousAdjustRequest
	GetIris() *string
	SetOwnerId(v int64) *ContinuousAdjustRequest
	GetOwnerId() *int64
}

type ContinuousAdjustRequest struct {
	// example:
	//
	// 0.5
	Focus *string `json:"Focus,omitempty" xml:"Focus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0.5
	Iris    *string `json:"Iris,omitempty" xml:"Iris,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ContinuousAdjustRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinuousAdjustRequest) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustRequest) GetFocus() *string {
	return s.Focus
}

func (s *ContinuousAdjustRequest) GetId() *string {
	return s.Id
}

func (s *ContinuousAdjustRequest) GetIris() *string {
	return s.Iris
}

func (s *ContinuousAdjustRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ContinuousAdjustRequest) SetFocus(v string) *ContinuousAdjustRequest {
	s.Focus = &v
	return s
}

func (s *ContinuousAdjustRequest) SetId(v string) *ContinuousAdjustRequest {
	s.Id = &v
	return s
}

func (s *ContinuousAdjustRequest) SetIris(v string) *ContinuousAdjustRequest {
	s.Iris = &v
	return s
}

func (s *ContinuousAdjustRequest) SetOwnerId(v int64) *ContinuousAdjustRequest {
	s.OwnerId = &v
	return s
}

func (s *ContinuousAdjustRequest) Validate() error {
	return dara.Validate(s)
}
