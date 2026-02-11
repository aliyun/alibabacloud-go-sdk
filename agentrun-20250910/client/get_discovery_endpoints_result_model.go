// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveryEndpointsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDiscoveryEndpointsResult
	GetCode() *string
	SetData(v *GetDiscoveryEndpointsOutput) *GetDiscoveryEndpointsResult
	GetData() *GetDiscoveryEndpointsOutput
	SetRequestId(v string) *GetDiscoveryEndpointsResult
	GetRequestId() *string
}

type GetDiscoveryEndpointsResult struct {
	Code      *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data      *GetDiscoveryEndpointsOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDiscoveryEndpointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveryEndpointsResult) GoString() string {
	return s.String()
}

func (s *GetDiscoveryEndpointsResult) GetCode() *string {
	return s.Code
}

func (s *GetDiscoveryEndpointsResult) GetData() *GetDiscoveryEndpointsOutput {
	return s.Data
}

func (s *GetDiscoveryEndpointsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiscoveryEndpointsResult) SetCode(v string) *GetDiscoveryEndpointsResult {
	s.Code = &v
	return s
}

func (s *GetDiscoveryEndpointsResult) SetData(v *GetDiscoveryEndpointsOutput) *GetDiscoveryEndpointsResult {
	s.Data = v
	return s
}

func (s *GetDiscoveryEndpointsResult) SetRequestId(v string) *GetDiscoveryEndpointsResult {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveryEndpointsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
