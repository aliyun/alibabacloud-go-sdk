// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRandomPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeCharacters(v string) *GetRandomPasswordRequest
	GetExcludeCharacters() *string
	SetExcludeLowercase(v string) *GetRandomPasswordRequest
	GetExcludeLowercase() *string
	SetExcludeNumbers(v string) *GetRandomPasswordRequest
	GetExcludeNumbers() *string
	SetExcludePunctuation(v string) *GetRandomPasswordRequest
	GetExcludePunctuation() *string
	SetExcludeUppercase(v string) *GetRandomPasswordRequest
	GetExcludeUppercase() *string
	SetPasswordLength(v string) *GetRandomPasswordRequest
	GetPasswordLength() *string
	SetRequireEachIncludedType(v string) *GetRandomPasswordRequest
	GetRequireEachIncludedType() *string
}

type GetRandomPasswordRequest struct {
	// The characters that are not included in the password to be generated.
	//
	// Valid values:
	//
	// ` Valid characters: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ! \\"#$%&\\"()*+,-. /:;<=>? @[\\] your_project_id} ~  `.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// ABCabc
	ExcludeCharacters *string `json:"ExcludeCharacters,omitempty" xml:"ExcludeCharacters,omitempty"`
	// Specifies whether to exclude lowercase letters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeLowercase *string `json:"ExcludeLowercase,omitempty" xml:"ExcludeLowercase,omitempty"`
	// Specifies whether to exclude digits.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeNumbers *string `json:"ExcludeNumbers,omitempty" xml:"ExcludeNumbers,omitempty"`
	// Specifies whether to exclude special characters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludePunctuation *string `json:"ExcludePunctuation,omitempty" xml:"ExcludePunctuation,omitempty"`
	// Specifies whether to exclude uppercase letters.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ExcludeUppercase *string `json:"ExcludeUppercase,omitempty" xml:"ExcludeUppercase,omitempty"`
	// The number of bytes that the password to be generated contains.
	//
	// Valid values: 8 to 128.
	//
	// Default value: 32
	//
	// example:
	//
	// 32
	PasswordLength *string `json:"PasswordLength,omitempty" xml:"PasswordLength,omitempty"`
	// Specifies whether to include all the preceding character types.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	RequireEachIncludedType *string `json:"RequireEachIncludedType,omitempty" xml:"RequireEachIncludedType,omitempty"`
}

func (s GetRandomPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRandomPasswordRequest) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordRequest) GetExcludeCharacters() *string {
	return s.ExcludeCharacters
}

func (s *GetRandomPasswordRequest) GetExcludeLowercase() *string {
	return s.ExcludeLowercase
}

func (s *GetRandomPasswordRequest) GetExcludeNumbers() *string {
	return s.ExcludeNumbers
}

func (s *GetRandomPasswordRequest) GetExcludePunctuation() *string {
	return s.ExcludePunctuation
}

func (s *GetRandomPasswordRequest) GetExcludeUppercase() *string {
	return s.ExcludeUppercase
}

func (s *GetRandomPasswordRequest) GetPasswordLength() *string {
	return s.PasswordLength
}

func (s *GetRandomPasswordRequest) GetRequireEachIncludedType() *string {
	return s.RequireEachIncludedType
}

func (s *GetRandomPasswordRequest) SetExcludeCharacters(v string) *GetRandomPasswordRequest {
	s.ExcludeCharacters = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeLowercase(v string) *GetRandomPasswordRequest {
	s.ExcludeLowercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeNumbers(v string) *GetRandomPasswordRequest {
	s.ExcludeNumbers = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludePunctuation(v string) *GetRandomPasswordRequest {
	s.ExcludePunctuation = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeUppercase(v string) *GetRandomPasswordRequest {
	s.ExcludeUppercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetPasswordLength(v string) *GetRandomPasswordRequest {
	s.PasswordLength = &v
	return s
}

func (s *GetRandomPasswordRequest) SetRequireEachIncludedType(v string) *GetRandomPasswordRequest {
	s.RequireEachIncludedType = &v
	return s
}

func (s *GetRandomPasswordRequest) Validate() error {
	return dara.Validate(s)
}
