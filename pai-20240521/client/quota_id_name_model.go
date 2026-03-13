// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaIdName interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *QuotaIdName
	GetQuotaId() *string
	SetQuotaName(v string) *QuotaIdName
	GetQuotaName() *string
}

type QuotaIdName struct {
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// dlc-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s QuotaIdName) String() string {
	return dara.Prettify(s)
}

func (s QuotaIdName) GoString() string {
	return s.String()
}

func (s *QuotaIdName) GetQuotaId() *string {
	return s.QuotaId
}

func (s *QuotaIdName) GetQuotaName() *string {
	return s.QuotaName
}

func (s *QuotaIdName) SetQuotaId(v string) *QuotaIdName {
	s.QuotaId = &v
	return s
}

func (s *QuotaIdName) SetQuotaName(v string) *QuotaIdName {
	s.QuotaName = &v
	return s
}

func (s *QuotaIdName) Validate() error {
	return dara.Validate(s)
}
