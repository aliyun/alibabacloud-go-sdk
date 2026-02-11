// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeShareUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetRetcodeShareUrlRequest
	GetPid() *string
}

type GetRetcodeShareUrlRequest struct {
	// This parameter is required.
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetRetcodeShareUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeShareUrlRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeShareUrlRequest) SetPid(v string) *GetRetcodeShareUrlRequest {
	s.Pid = &v
	return s
}

func (s *GetRetcodeShareUrlRequest) Validate() error {
	return dara.Validate(s)
}
