// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpListEnvInfoResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*EnvInfoDO) *PdpListEnvInfoResult
	GetData() []*EnvInfoDO
	SetRequestId(v string) *PdpListEnvInfoResult
	GetRequestId() *string
}

type PdpListEnvInfoResult struct {
	Data []*EnvInfoDO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 121
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PdpListEnvInfoResult) String() string {
	return dara.Prettify(s)
}

func (s PdpListEnvInfoResult) GoString() string {
	return s.String()
}

func (s *PdpListEnvInfoResult) GetData() []*EnvInfoDO {
	return s.Data
}

func (s *PdpListEnvInfoResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpListEnvInfoResult) SetData(v []*EnvInfoDO) *PdpListEnvInfoResult {
	s.Data = v
	return s
}

func (s *PdpListEnvInfoResult) SetRequestId(v string) *PdpListEnvInfoResult {
	s.RequestId = &v
	return s
}

func (s *PdpListEnvInfoResult) Validate() error {
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
