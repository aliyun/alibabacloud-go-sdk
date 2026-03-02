// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PdpCommonResult
	GetData() *bool
	SetRequestId(v string) *PdpCommonResult
	GetRequestId() *string
}

type PdpCommonResult struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PdpCommonResult) String() string {
	return dara.Prettify(s)
}

func (s PdpCommonResult) GoString() string {
	return s.String()
}

func (s *PdpCommonResult) GetData() *bool {
	return s.Data
}

func (s *PdpCommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpCommonResult) SetData(v bool) *PdpCommonResult {
	s.Data = &v
	return s
}

func (s *PdpCommonResult) SetRequestId(v string) *PdpCommonResult {
	s.RequestId = &v
	return s
}

func (s *PdpCommonResult) Validate() error {
	return dara.Validate(s)
}
