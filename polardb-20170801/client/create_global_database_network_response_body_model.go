// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGDNId(v string) *CreateGlobalDatabaseNetworkResponseBody
	GetGDNId() *string
	SetRequestId(v string) *CreateGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
}

type CreateGlobalDatabaseNetworkResponseBody struct {
	// The ID of the GDN.
	//
	// example:
	//
	// gd-m5ex5wqfqbou*****
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkResponseBody) GetGDNId() *string {
	return s.GDNId
}

func (s *CreateGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalDatabaseNetworkResponseBody) SetGDNId(v string) *CreateGlobalDatabaseNetworkResponseBody {
	s.GDNId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *CreateGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
