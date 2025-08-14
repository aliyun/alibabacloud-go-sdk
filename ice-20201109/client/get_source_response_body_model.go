// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSourceResponseBody
	GetRequestId() *string
	SetSource(v *ChannelAssemblySource) *GetSourceResponseBody
	GetSource() *ChannelAssemblySource
}

type GetSourceResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source information.
	Source *ChannelAssemblySource `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceResponseBody) GetSource() *ChannelAssemblySource {
	return s.Source
}

func (s *GetSourceResponseBody) SetRequestId(v string) *GetSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceResponseBody) SetSource(v *ChannelAssemblySource) *GetSourceResponseBody {
	s.Source = v
	return s
}

func (s *GetSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
