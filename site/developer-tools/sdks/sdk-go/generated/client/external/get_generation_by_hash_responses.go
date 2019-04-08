// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetGenerationByHashReader is a Reader for the GetGenerationByHash structure.
type GetGenerationByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGenerationByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGenerationByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetGenerationByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetGenerationByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGenerationByHashOK creates a GetGenerationByHashOK with default headers values
func NewGetGenerationByHashOK() *GetGenerationByHashOK {
	return &GetGenerationByHashOK{}
}

/*GetGenerationByHashOK handles this case with default header values.

Successful operation
*/
type GetGenerationByHashOK struct {
	Payload *models.Generation
}

func (o *GetGenerationByHashOK) Error() string {
	return fmt.Sprintf("[GET /generations/hash/{hash}][%d] getGenerationByHashOK  %+v", 200, o.Payload)
}

func (o *GetGenerationByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Generation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGenerationByHashBadRequest creates a GetGenerationByHashBadRequest with default headers values
func NewGetGenerationByHashBadRequest() *GetGenerationByHashBadRequest {
	return &GetGenerationByHashBadRequest{}
}

/*GetGenerationByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetGenerationByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetGenerationByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /generations/hash/{hash}][%d] getGenerationByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetGenerationByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGenerationByHashNotFound creates a GetGenerationByHashNotFound with default headers values
func NewGetGenerationByHashNotFound() *GetGenerationByHashNotFound {
	return &GetGenerationByHashNotFound{}
}

/*GetGenerationByHashNotFound handles this case with default header values.

Generation not found
*/
type GetGenerationByHashNotFound struct {
	Payload *models.Error
}

func (o *GetGenerationByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /generations/hash/{hash}][%d] getGenerationByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetGenerationByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
