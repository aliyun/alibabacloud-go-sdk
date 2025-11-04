// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSourceLocationResponseBody
	GetRequestId() *string
	SetSourceLocation(v *ChannelAssemblySourceLocation) *GetSourceLocationResponseBody
	GetSourceLocation() *ChannelAssemblySourceLocation
}

type GetSourceLocationResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source location information.
	SourceLocation *ChannelAssemblySourceLocation `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s GetSourceLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceLocationResponseBody) GetSourceLocation() *ChannelAssemblySourceLocation {
	return s.SourceLocation
}

func (s *GetSourceLocationResponseBody) SetRequestId(v string) *GetSourceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceLocationResponseBody) SetSourceLocation(v *ChannelAssemblySourceLocation) *GetSourceLocationResponseBody {
	s.SourceLocation = v
	return s
}

func (s *GetSourceLocationResponseBody) Validate() error {
	if s.SourceLocation != nil {
		if err := s.SourceLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
