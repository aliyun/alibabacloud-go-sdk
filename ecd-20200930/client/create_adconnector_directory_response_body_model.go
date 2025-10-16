// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdConnectors(v []*CreateADConnectorDirectoryResponseBodyAdConnectors) *CreateADConnectorDirectoryResponseBody
	GetAdConnectors() []*CreateADConnectorDirectoryResponseBodyAdConnectors
	SetDirectoryId(v string) *CreateADConnectorDirectoryResponseBody
	GetDirectoryId() *string
	SetRequestId(v string) *CreateADConnectorDirectoryResponseBody
	GetRequestId() *string
	SetTrustPassword(v string) *CreateADConnectorDirectoryResponseBody
	GetTrustPassword() *string
}

type CreateADConnectorDirectoryResponseBody struct {
	// The details of AD connectors.
	AdConnectors []*CreateADConnectorDirectoryResponseBodyAdConnectors `json:"AdConnectors,omitempty" xml:"AdConnectors,omitempty" type:"Repeated"`
	// The ID of the AD directory.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3FE99D5E-93A1-493F-B1CB-0ABD4D05BEFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The AD trust password.
	//
	// example:
	//
	// 82Tg****
	TrustPassword *string `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
}

func (s CreateADConnectorDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBody) GetAdConnectors() []*CreateADConnectorDirectoryResponseBodyAdConnectors {
	return s.AdConnectors
}

func (s *CreateADConnectorDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateADConnectorDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateADConnectorDirectoryResponseBody) GetTrustPassword() *string {
	return s.TrustPassword
}

func (s *CreateADConnectorDirectoryResponseBody) SetAdConnectors(v []*CreateADConnectorDirectoryResponseBodyAdConnectors) *CreateADConnectorDirectoryResponseBody {
	s.AdConnectors = v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetDirectoryId(v string) *CreateADConnectorDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetRequestId(v string) *CreateADConnectorDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetTrustPassword(v string) *CreateADConnectorDirectoryResponseBody {
	s.TrustPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) Validate() error {
	if s.AdConnectors != nil {
		for _, item := range s.AdConnectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateADConnectorDirectoryResponseBodyAdConnectors struct {
	// The connection address.
	//
	// example:
	//
	// ``127.0.**.**``
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBodyAdConnectors) GetAddress() *string {
	return s.Address
}

func (s *CreateADConnectorDirectoryResponseBodyAdConnectors) SetAddress(v string) *CreateADConnectorDirectoryResponseBodyAdConnectors {
	s.Address = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBodyAdConnectors) Validate() error {
	return dara.Validate(s)
}
