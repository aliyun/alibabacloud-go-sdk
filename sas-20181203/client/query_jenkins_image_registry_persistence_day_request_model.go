// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJenkinsImageRegistryPersistenceDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *QueryJenkinsImageRegistryPersistenceDayRequest
	GetSourceIp() *string
}

type QueryJenkinsImageRegistryPersistenceDayRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 113.83.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s QueryJenkinsImageRegistryPersistenceDayRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJenkinsImageRegistryPersistenceDayRequest) GoString() string {
	return s.String()
}

func (s *QueryJenkinsImageRegistryPersistenceDayRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *QueryJenkinsImageRegistryPersistenceDayRequest) SetSourceIp(v string) *QueryJenkinsImageRegistryPersistenceDayRequest {
	s.SourceIp = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayRequest) Validate() error {
	return dara.Validate(s)
}
