// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpSlsLogInfo interface {
	dara.Model
	String() string
	GoString() string
	SetSlsUrl(v string) *PdpSlsLogInfo
	GetSlsUrl() *string
}

type PdpSlsLogInfo struct {
	// example:
	//
	// https://sls.console.aliyun.com/lognext/project/k8s-log-c42539518786e49fbb390929599dec9fa/logsearch/yunmall-customer112production
	SlsUrl *string `json:"slsUrl,omitempty" xml:"slsUrl,omitempty"`
}

func (s PdpSlsLogInfo) String() string {
	return dara.Prettify(s)
}

func (s PdpSlsLogInfo) GoString() string {
	return s.String()
}

func (s *PdpSlsLogInfo) GetSlsUrl() *string {
	return s.SlsUrl
}

func (s *PdpSlsLogInfo) SetSlsUrl(v string) *PdpSlsLogInfo {
	s.SlsUrl = &v
	return s
}

func (s *PdpSlsLogInfo) Validate() error {
	return dara.Validate(s)
}
