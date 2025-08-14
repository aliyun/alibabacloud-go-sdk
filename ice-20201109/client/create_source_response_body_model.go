// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSourceResponseBody
	GetRequestId() *string
	SetSource(v *ChannelAssemblySource) *CreateSourceResponseBody
	GetSource() *ChannelAssemblySource
}

type CreateSourceResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source information.
	Source *ChannelAssemblySource `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSourceResponseBody) GetSource() *ChannelAssemblySource {
	return s.Source
}

func (s *CreateSourceResponseBody) SetRequestId(v string) *CreateSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSourceResponseBody) SetSource(v *ChannelAssemblySource) *CreateSourceResponseBody {
	s.Source = v
	return s
}

func (s *CreateSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
