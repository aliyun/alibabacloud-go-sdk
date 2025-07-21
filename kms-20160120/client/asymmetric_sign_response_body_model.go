// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *AsymmetricSignResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricSignResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *AsymmetricSignResponseBody
	GetRequestId() *string
	SetValue(v string) *AsymmetricSignResponseBody
	GetValue() *string
}

type AsymmetricSignResponseBody struct {
	// The version ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The digest that is generated for the original message by using a hash algorithm. The hash algorithm is specified by the Algorithm parameter.
	//
	// > 	- The value is encoded in Base64.
	//
	// > 	- For more information about how to calculate message digests, see the **Preprocess signature: compute a message digest*	- section of the [Generate and verify a signature by using an asymmetric CMK](https://help.aliyun.com/document_detail/148146.html) topic.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The calculated signature.
	//
	// >  The value is encoded in Base64.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricSignResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricSignResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsymmetricSignResponseBody) GetValue() *string {
	return s.Value
}

func (s *AsymmetricSignResponseBody) SetKeyId(v string) *AsymmetricSignResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetKeyVersionId(v string) *AsymmetricSignResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetRequestId(v string) *AsymmetricSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricSignResponseBody) SetValue(v string) *AsymmetricSignResponseBody {
	s.Value = &v
	return s
}

func (s *AsymmetricSignResponseBody) Validate() error {
	return dara.Validate(s)
}
